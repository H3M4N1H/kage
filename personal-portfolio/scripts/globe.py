"""Generate the local point-cloud globe from Natural Earth public-domain land data."""
import json, math
from pathlib import Path
root=Path(__file__).resolve().parents[1]
data=json.loads((root/'scripts/data/land.geojson').read_text())
polygons=[]
for f in data['features']:
 g=f['geometry']; ps=g['coordinates'] if g['type']=='MultiPolygon' else [g['coordinates']]
 for rings in ps:
  ring=rings[0];xs=[p[0] for p in ring];ys=[p[1] for p in ring]
  polygons.append((min(xs),max(xs),min(ys),max(ys),ring))
def land(x,y):
 for x0,x1,y0,y1,ring in polygons:
  if not(x0<=x<=x1 and y0<=y<=y1):continue
  inside=False;j=len(ring)-1
  for i in range(len(ring)):
   xi,yi=ring[i];xj,yj=ring[j]
   if (yi>y)!=(yj>y) and x<(xj-xi)*(y-yi)/(yj-yi)+xi:inside=not inside
   j=i
  if inside:return True
 return False
lat0=math.radians(16);lon0=math.radians(68);r=213

def project(lon,lat):
 a=math.radians(lon)-lon0;b=math.radians(lat)
 z=math.sin(lat0)*math.sin(b)+math.cos(lat0)*math.cos(b)*math.cos(a)
 return 270+r*math.cos(b)*math.sin(a),252-r*(math.cos(lat0)*math.sin(b)-math.sin(lat0)*math.cos(b)*math.cos(a)),z
svg=['<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 540 510" fill="none"><defs><radialGradient id="g"><stop stop-color="#aaff24" stop-opacity=".05"/><stop offset="1" stop-color="#050807" stop-opacity="0"/></radialGradient></defs><circle cx="270" cy="252" r="245" fill="url(#g)"/><circle cx="270" cy="252" r="213" stroke="#7d9271" stroke-opacity=".22"/>']
# Sparse latitude and longitude graticules.
for axis in ['lat','lon']:
 for fixed in range(-150,181,30):
  if axis=='lat' and abs(fixed)>75:continue
  path=[]
  for t in range(-180 if axis=='lat' else -90,181 if axis=='lat' else 91,2):
   x,y,z=project(t,fixed) if axis=='lat' else project(fixed,t)
   if z>0:path.append(('M' if not path or path[-1]=='|' else 'L')+f'{x:.1f},{y:.1f}')
   elif path and path[-1]!='|':path.append('|')
  svg.append('<path d="'+''.join(path).replace('|','')+'" stroke="#53644b" stroke-opacity=".18" stroke-width=".6"/>')
for lat in range(-84,86,2):
 step=1.65/max(.2,math.cos(math.radians(lat)))
 lon=-180
 while lon<180:
  x,y,z=project(lon,lat)
  if z>.02:
   onland=land(lon,lat)
   if onland:
    highlight=(56<lon<92 and 7<lat<33) or (-10<lon<35 and 48<lat<65)
    color='#a6ed42' if highlight else '#b2beb5'
    opacity=(.36+.47*z) if highlight else (.22+.46*z)
    svg.append(f'<circle cx="{x:.2f}" cy="{y:.2f}" r="{.68+.32*z:.2f}" fill="{color}" opacity="{opacity:.2f}"/>')
   elif int(lon*100)%4==0:
    svg.append(f'<circle cx="{x:.2f}" cy="{y:.2f}" r=".5" fill="#70816e" opacity=".18"/>')
  lon+=step
svg.append('<ellipse cx="270" cy="252" rx="262" ry="81" transform="rotate(-49 270 252)" stroke="#a7f32b" stroke-width=".7" stroke-opacity=".43"/><ellipse cx="270" cy="252" rx="228" ry="94" transform="rotate(31 270 252)" stroke="#a7f32b" stroke-width=".6" stroke-opacity=".18" stroke-dasharray="2 8"/><path d="M40 68h26m-13-13v26M464 441h26m-13-13v26" stroke="#687067" stroke-opacity=".5"/>')
x,y,_=project(80.27,13.08)
svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="12" stroke="#b0ff23"/><circle cx="{x:.1f}" cy="{y:.1f}" r="3" fill="#b0ff23"/><path d="M{x-18:.1f} {y:.1f}h36M{x:.1f} {y-18:.1f}v36M{x+10:.1f} {y+10:.1f}l22 22h23" stroke="#b0ff23"/><rect x="{x+51:.1f}" y="{y+18:.1f}" width="122" height="41" fill="#060a07" fill-opacity=".85"/><text x="{x+60:.1f}" y="{y+33:.1f}" fill="#b0ff23" font-family="monospace" font-size="12" font-weight="bold">CHENNAI, IN</text><text x="{x+60:.1f}" y="{y+50:.1f}" fill="#a2ada0" font-family="monospace" font-size="10">Based in India</text></svg>')
(root/'public/globe.svg').write_text(''.join(svg))
print('Generated globe.svg')
points=[]
for lat in range(-84,86):
 step=1.25/max(.2,math.cos(math.radians(lat)))
 lon=-180
 while lon<180:
  if land(lon,lat):points.append([round(lon,2),lat])
  lon+=step
(root/'public/globe-points.json').write_text(json.dumps(points,separators=(',',':')))
