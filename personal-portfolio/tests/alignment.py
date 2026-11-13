from playwright.sync_api import sync_playwright
import json, os
from pathlib import Path
content=json.loads((Path(__file__).parents[1]/"src/content.json").read_text())
project_count=len(content["projects"])
systems_count=sum(p["category"]=="Systems & tooling" for p in content["projects"])
security_count=sum(p["category"]=="Security & learning" for p in content["projects"])
with sync_playwright() as p:
 b=p.chromium.launch(headless=True,executable_path=os.environ.get('CHROMIUM_PATH'))
 page=b.new_page(reduced_motion='reduce');errors=[];page.on('pageerror',lambda e:errors.append(str(e)));results=[]
 for w,h in [(320,740),(360,800),(390,844),(600,900),(700,900),(768,1024),(900,900),(1024,768),(1180,820),(1280,720),(1366,768),(1440,900),(1920,1080),(2560,1440)]:
  page.set_viewport_size({'width':w,'height':h});page.goto('http://127.0.0.1:4173');page.evaluate('document.fonts.ready');page.locator('portfolio-globe[data-ready=true]').wait_for()
  assert page.evaluate('document.documentElement.scrollWidth<=innerWidth'),w
  gutter=16 if w<=600 else 24 if w<=900 else min(64,max(32,w*.04))
  edge=max(0,(w-1440)/2)+gutter
  layout=page.evaluate('''()=>Object.fromEntries(['.hero','.about-section','.dashboard','#work','#events','#experience','#contact'].map(s=>{const b=document.querySelector(s).getBoundingClientRect();return [s,{left:b.left,right:b.right,top:b.top,bottom:b.bottom}]}))''')
  for name,r in layout.items():
   assert abs(r['left']-edge)<1,(w,name,r)
   assert abs(r['right']-(w-edge))<1,(w,name,r)
  for selector in ['.header-inner','footer']:
   rect=page.locator(selector).bounding_box()
   inset=float(page.locator(selector).evaluate('e=>getComputedStyle(e).paddingLeft').replace('px',''))
   assert abs(rect['x']+inset-edge)<1,(w,selector,rect)
  if w<=600:
   hero=page.locator('.hero').bounding_box()
   for selector in ['.hero-copy','.globe-window','.terminal']:
    rect=page.locator(selector).bounding_box()
    assert abs(rect['x']-(edge+8))<1,(w,selector,rect)
    assert abs(rect['x']+rect['width']-(w-edge-8))<1,(w,selector,rect)
   actions=page.locator('.hero-actions .button').all()
   for button in actions: assert button.bounding_box()['height']>=48
   if w<=380: assert actions[1].bounding_box()['y']>=actions[0].bounding_box()['y']+48
  names=list(layout)
  for prev,nxt in zip(names,names[1:]):assert layout[prev]['bottom']<=layout[nxt]['top']+1,(w,prev,nxt)
  assert page.locator('article.project:visible').count()==4
  assert page.locator('.project-repo').count()==4
  assert all(a.get_attribute('href').startswith('https://github.com/H3M4N1H/') for a in page.locator('.project-repo').all())
  for summary in page.locator('.event-history summary').all(): summary.click()
  assert page.locator('.event-history details[open]').count()==2
  clips=page.evaluate('''()=>Array.from(document.querySelectorAll('h1,h2,h3,h4,p,button,li,small')).filter(e=>e.getClientRects().length&&!e.closest('[aria-hidden="true"],.sr-only')).filter(e=>e.clientWidth>0&&e.scrollWidth>e.clientWidth+2).map(e=>({tag:e.tagName,text:e.textContent.slice(0,50),width:e.clientWidth,scroll:e.scrollWidth}))''')
  assert not clips,(w,clips)
  for summary in page.locator('.event-history summary').all(): summary.click()
  positions={}
  for target in ['about','skills','work','events','experience']:
   if w<=900:page.locator('.menu-toggle').click()
   page.locator(f'nav a[href="#{target}"]').click()
   rect=page.locator('#'+target).bounding_box();header=page.locator('.site-header').bounding_box()
   assert rect['y']>=header['height']-1,(w,target,rect['y'],header['height'])
   assert rect['y']<h,(w,target,'anchor below viewport')
   positions[target]=page.evaluate('scrollY')
  assert abs(positions['about']-positions['skills'])>100,(w,positions)
  for name in ['Development','Tools','Systems','Security']:
   page.get_by_role('tab',name=name,exact=True).click()
   assert page.locator('.skill-grid:visible .skill-card').count()==8
  page.get_by_role('button',name='Systems & tooling',exact=True).click();assert page.locator('article.project:visible').count()==systems_count
  page.get_by_role('button',name='All work').click();assert page.locator('article.project:visible').count()==project_count
  page.evaluate('scrollTo(0,0)')
  if w in [360,1366,1920]:page.screenshot(path=f'/tmp/portfolio-aligned-{w}.png',full_page=True)
  results.append({'width':w,'gutter':gutter,'about_skills_separation':round(abs(positions['about']-positions['skills']))})
 assert not errors,errors
 print(json.dumps(results));print('All section edges, anchor visibility, About/Skills separation, text clipping, filters, tabs, and overflow checks passed.')
 b.close()
