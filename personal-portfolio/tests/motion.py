from playwright.sync_api import sync_playwright
import os
with sync_playwright() as p:
 b=p.chromium.launch(headless=True,executable_path=os.environ.get('CHROMIUM_PATH'))
 page=b.new_page(viewport={'width':1440,'height':900});errors=[];page.on('pageerror',lambda e:errors.append(str(e)))
 page.add_init_script('''window.reveals=[];const original=Element.prototype.animate;Element.prototype.animate=function(frames,options){window.reveals.push({element:this.className||this.id,duration:options.duration,transform:frames[0].transform,opacity:frames[0].opacity});return original.call(this,frames,options)}''')
 page.goto('http://127.0.0.1:4173');page.wait_for_function('window.reveals.length>0')
 assert page.evaluate('reveals.every(r=>r.duration===900 && r.opacity===undefined)')
 page.wait_for_function("document.querySelector('.hero-copy').getAnimations().length===0")
 page.locator('#work .section-heading').evaluate("e=>e.scrollIntoView({block:'center',behavior:'instant'})")
 page.wait_for_function("reveals.some(r=>r.element==='section-heading')")
 page.wait_for_function("document.querySelector('#work .section-heading').getAnimations().length===0")
 page.wait_for_timeout(1100)
 assert page.evaluate("reveals.some(r=>r.transform.startsWith('translateX(-')) && reveals.some(r=>r.transform.startsWith('translateX(32')) && reveals.some(r=>r.transform.startsWith('translateY('))")
 count=page.evaluate('reveals.length')
 page.evaluate('scrollTo({top:0,behavior:"instant"})');page.wait_for_timeout(1100)
 page.locator('#work .section-heading').evaluate("e=>e.scrollIntoView({block:'center',behavior:'instant'})");page.wait_for_timeout(1100)
 assert page.evaluate('reveals.length')==count
 page.emulate_media(reduced_motion='reduce');page.locator('#contact').evaluate("e=>e.scrollIntoView({behavior:'instant'})");page.wait_for_timeout(1100)
 assert page.evaluate('reveals.length')==count
 # Sample rendered frames during a real entrance and focus an unrelated control.
 stable=b.new_page(viewport={'width':1440,'height':900})
 stable.goto('http://127.0.0.1:4173')
 stable.evaluate("document.querySelector('#work').scrollIntoView({behavior:'instant'})")
 stable.wait_for_function("document.querySelector('.project').getAnimations().length>0")
 assert stable.evaluate("""async () => {
   const card=document.querySelector('.project');
   const animation=card.getAnimations()[0];
   document.querySelector('.menu-toggle').dispatchEvent(new FocusEvent('focusin',{bubbles:true}));
   if(animation.playState==='idle') return false;
   for(let i=0;i<20;i++){
     await new Promise(requestAnimationFrame);
     if(getComputedStyle(card).opacity!=='1') return false;
     if(document.documentElement.scrollWidth>innerWidth) return false;
   }
   return true;
 }""")
 stable.close()
 assert not errors,errors
 reduced=b.new_page(reduced_motion='reduce',viewport={'width':390,'height':844});reduced.goto('http://127.0.0.1:4173');reduced.locator('portfolio-globe[data-ready=true]').wait_for()
 assert reduced.locator('.hero-copy').evaluate('e=>e.getAnimations().length')==0
 assert reduced.locator('.terminal').bounding_box()['y']>=864
 assert reduced.evaluate('document.documentElement.scrollWidth<=innerWidth')
 print('No opacity resets; directional 900ms one-time reveals, repeat-scroll behavior, reduced-motion preference, and mobile opening-screen layout passed.')
 b.close()
