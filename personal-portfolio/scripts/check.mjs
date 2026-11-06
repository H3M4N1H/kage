import assert from 'node:assert/strict';
import {readFile} from 'node:fs/promises';
const c = JSON.parse(await readFile('src/content.json','utf8'));
assert.equal(c.profiles.length,4);
for(const p of c.profiles) assert.equal(new URL(p.url).protocol,'https:');
assert.equal(new Set(c.projects.map(p=>p.name)).size,c.projects.length);
for(const p of c.projects) for(const key of ['name','category','summary','description','visual','status']) assert.ok(p[key],`${p.name}: ${key}`);
for(const r of [...c.results,...c.roles]) assert.equal(typeof r.confirmed,'boolean');
console.log('Content validation passed.');
