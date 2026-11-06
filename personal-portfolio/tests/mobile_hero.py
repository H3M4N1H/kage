from playwright.sync_api import sync_playwright
import os
with sync_playwright() as p:
 b=p.chromium.launch(headless=True,executable_path=os.environ.get('CHROMIUM_PATH'))
 page=b.new_page(is_mobile=True,has_touch=True,reduced_motion='reduce')
 for w,h in [(320,568),(360,800),(375,667),(390,844),(412,915),(430,932),(600,900)]:
  page.set_viewport_size({'width':w,'height':h});page.goto('http://127.0.0.1:4173');page.evaluate('document.fonts.ready');page.locator('portfolio-globe[data-ready=true]').wait_for()
  assert page.locator('.terminal').bounding_box()['y']>=h+20,(w,h,page.locator('.terminal').bounding_box())
  assert page.evaluate('document.documentElement.scrollWidth<=innerWidth')
  intro=page.locator('.hero-copy').bounding_box();globe=page.locator('.globe-window').bounding_box()
  assert intro['y']+intro['height']<=globe['y']
  if w==390:page.screenshot(path='/tmp/mobile-first-screen.png')
 print('Profile box starts below the first screen at seven phone sizes; no overlap or horizontal overflow.')
 b.close()
