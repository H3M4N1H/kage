from playwright.sync_api import sync_playwright
import os

with sync_playwright() as p:
 b=p.chromium.launch(headless=True,executable_path=os.environ.get('CHROMIUM_PATH'))
 page=b.new_page(viewport={'width':1440,'height':900}); errors=[]
 page.on('pageerror',lambda e:errors.append(str(e)))
 page.goto('http://127.0.0.1:4173')
 assert page.locator('html').evaluate("e=>e.classList.contains('motion-ready')")
 # Hidden before scrolling, not painted and then reset when the observer fires.
 card=page.locator('.project').first
 assert card.evaluate("e=>getComputedStyle(e).opacity")=='0'
 assert page.locator('.hero-actions').evaluate("e=>getComputedStyle(e).animationDelay")=='0.34s'
 assert page.locator('.featured-panel').evaluate("e=>getComputedStyle(e).transitionDelay")=='0.12s'
 assert card.evaluate("e=>getComputedStyle(e).transitionDuration")=='0.6s'
 card.evaluate("e=>e.scrollIntoView({behavior:'instant',block:'center'})")
 page.wait_for_function("document.querySelector('.project').classList.contains('is-visible')")
 assert page.evaluate("""async()=>{
  const card=document.querySelector('.project');let previous=0;
  for(let i=0;i<65;i++){
   await new Promise(requestAnimationFrame);
   const value=Number(getComputedStyle(card).opacity);
   if(value+0.001<previous || document.documentElement.scrollWidth>innerWidth)return false;
   previous=value;
  }
  return previous===1;
 }""")
 page.evaluate("scrollTo({top:0,behavior:'instant'})")
 card.evaluate("e=>e.scrollIntoView({behavior:'instant',block:'center'})")
 assert card.evaluate("e=>getComputedStyle(e).opacity")=='1'
 # Keyboard navigation reveals a pending target without waiting on its fade.
 pending=page.locator('.profile-list')
 pending.locator('a').first.focus()
 assert pending.evaluate("e=>getComputedStyle(e).opacity")=='1'
 page.emulate_media(reduced_motion='reduce')
 assert page.locator('[data-reveal]').evaluate_all("es=>es.every(e=>getComputedStyle(e).opacity==='1')")
 for width,height in [(320,740),(390,844),(768,1024),(1440,900)]:
  mobile=b.new_page(viewport={'width':width,'height':height},reduced_motion='reduce')
  mobile.goto('http://127.0.0.1:4173')
  assert mobile.locator('.hero-actions').evaluate("e=>getComputedStyle(e).animationName")=='none'
  assert mobile.evaluate('document.documentElement.scrollWidth<=innerWidth')
  if width<=390:assert mobile.locator('.terminal').bounding_box()['y']>=height+20
  mobile.close()
 nojs=b.new_page(java_script_enabled=False);nojs.goto('http://127.0.0.1:4173')
 assert nojs.locator('[data-reveal]').evaluate_all("es=>es.every(e=>getComputedStyle(e).opacity==='1')")
 failed=b.new_page();failed.route('**/site.js',lambda route:route.abort());failed.goto('http://127.0.0.1:4173')
 failed.wait_for_function("!document.documentElement.classList.contains('motion-ready')",timeout=6000)
 assert failed.locator('.project').first.evaluate("e=>getComputedStyle(e).opacity")=='1'
 assert not errors,errors
 print('Reference timings, pre-paint hidden state, monotonic fade, one-time reveals, keyboard focus, mobile layout, reduced motion, no-JS and failed-script fallback passed.')
 b.close()
