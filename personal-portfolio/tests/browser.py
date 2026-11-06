from playwright.sync_api import sync_playwright
import json, os
from pathlib import Path
content=json.loads((Path(__file__).parents[1]/"src/content.json").read_text())
project_count=len(content["projects"])
systems_count=sum(p["category"]=="Systems & tooling" for p in content["projects"])
security_count=sum(p["category"]=="Security & learning" for p in content["projects"])
with sync_playwright() as p:
 browser=p.chromium.launch(headless=True, executable_path=os.environ.get('CHROMIUM_PATH'))
 page=browser.new_page()
 errors=[]
 page.on('pageerror',lambda e:errors.append(str(e)))
 for width in [360,768,1440,1672]:
  page.set_viewport_size({'width':width,'height':941})
  response=page.goto('http://127.0.0.1:4173')
  assert response.status==200
  page.locator('portfolio-globe[data-ready="true"]').wait_for()
  globe=page.locator('portfolio-globe')
  canvas=page.locator('portfolio-globe canvas')
  assert globe.get_attribute('data-spinning')=='true'
  canvas.click()
  assert globe.get_attribute('data-spinning')=='false'
  assert page.locator('.globe-controls,.summary-strip').count()==0
  canvas.press('Home')
  initial=globe.get_attribute('data-longitude')
  canvas.press('ArrowRight')
  assert globe.get_attribute('data-longitude')!=initial
  canvas.press('+')
  assert float(globe.get_attribute('data-zoom'))>1
  canvas.press('Home')
  assert globe.get_attribute('data-longitude')==initial
  canvas=page.locator('portfolio-globe canvas')
  canvas.focus();page.keyboard.press('ArrowLeft')
  assert globe.get_attribute('data-longitude')!=initial
  page.keyboard.press('Home')
  assert globe.get_attribute('data-longitude')==initial
  canvas.evaluate("e=>e.scrollIntoView({block:'center',behavior:'instant'})")
  box=canvas.bounding_box()
  page.mouse.move(box['x']+box['width']/2,box['y']+box['height']/2)
  page.mouse.down();page.mouse.move(box['x']+box['width']/2+40,box['y']+box['height']/2,steps=5);page.mouse.up()
  assert globe.get_attribute('data-longitude')!=initial
  assert globe.get_attribute('data-spinning')=='false'
  canvas.click()
  assert globe.get_attribute('data-spinning')=='true'
  canvas.press('Home')
  page.evaluate('document.fonts.ready')
  assert page.evaluate('document.documentElement.scrollWidth <= innerWidth'),width
  assert page.locator('article.project:visible').count()==9
  page.locator('.more-projects').click()
  assert page.locator('article.project:visible').count()==project_count
  if width>1000:
   assert page.locator('.social-icons').bounding_box()['y']+page.locator('.social-icons').bounding_box()['height'] < page.locator('.about-section').bounding_box()['y'], 'Hero overlaps About'
  for name in ['Development','Tools','Systems','Security']:
   page.get_by_role('tab',name=name,exact=True).click()
   assert page.locator('[role="tabpanel"]:visible .skill-card').count()==8
   assert page.get_by_role('tab',name=name,exact=True).get_attribute('aria-selected')=='true'
  page.get_by_role('tab',name='Security',exact=True).focus()
  page.keyboard.press('ArrowRight')
  assert page.get_by_role('tab',name='Development',exact=True).get_attribute('aria-selected')=='true'
  page.get_by_role('tab',name='Security',exact=True).click()
  page.get_by_role('button',name='Systems & tooling',exact=True).click()
  assert page.locator('article.project:visible').count()==systems_count
  page.get_by_role('button',name='Security & learning',exact=True).click()
  page.locator('.featured-card').first.click()
  assert page.locator('#project-aegis').is_visible()
  assert page.locator('article.project:visible').count()==project_count
  page.get_by_role('button',name='Security & learning',exact=True).click()
  assert page.locator('article.project:visible').count()==security_count
  page.get_by_role('button',name='All work').click()
  if width<901:
   page.locator('.menu-toggle').click()
   assert page.locator('#main-nav').is_visible()
   page.keyboard.press('Escape')
   assert not page.locator('#main-nav').is_visible()
  for link in page.locator('nav a').all():
   if width<901:page.locator('.menu-toggle').click()
   target=link.get_attribute('href')
   link.click()
   assert page.locator(target).count()==1
   if width<901:assert not page.locator('#main-nav').is_visible()
  page.goto('http://127.0.0.1:4173')
  page.evaluate('document.fonts.ready')
  page.screenshot(path=f'/tmp/portfolio-restored-{width}.png',full_page=True)
  if width==1672:page.screenshot(path='/tmp/portfolio-restored-desktop.png')
 assert not errors,errors
 canvas=page.locator('portfolio-globe canvas')
 canvas.click()
 assert page.locator('portfolio-globe').get_attribute('data-spinning')=='false'
 canvas.click()
 assert page.locator('portfolio-globe').get_attribute('data-spinning')=='true'
 before=page.locator('portfolio-globe').get_attribute('data-longitude')
 page.wait_for_function("document.querySelector('portfolio-globe').dataset.longitude !== " + json.dumps(before))
 canvas.press('Space')
 assert page.locator('portfolio-globe').get_attribute('data-spinning')=='false'
 canvas.press('Enter')
 assert page.locator('portfolio-globe').get_attribute('data-spinning')=='true'
 page.emulate_media(reduced_motion='reduce')
 assert page.evaluate('getComputedStyle(document.documentElement).scrollBehavior')=='auto'
 context=browser.new_context(java_script_enabled=False)
 nojs=context.new_page();nojs.goto('http://127.0.0.1:4173');assert nojs.locator('article.project').count()==project_count
 fallback=browser.new_page()
 fallback.route('**/globe-points.json',lambda route:route.abort())
 fallback.goto('http://127.0.0.1:4173')
 assert fallback.locator('.globe-fallback').is_visible()
 print(json.dumps({'widths':[360,768,1440,1672],'globe_drag_keyboard_zoom_reset_rotation':'passed','globe_fallback':'passed','overflow':False,'filters':'passed','skills_tabs_and_keyboard':'passed','project_deep_links':'passed','mobile_menu':'passed','anchors':'passed','reduced_motion':'passed','no_javascript_content':'passed','page_errors':errors}))
 browser.close()
