const CACHE_NAME = 'mariogeo-v1';
const ASSETS = [
  '/',
  '/static/index.html',
  'https://unpkg.com/htmx.org@1.9.10'
];

self.addEventListener('install', (e) => {
  e.waitUntil(caches.open(CACHE_NAME).then(cache => cache.addAll(ASSETS)));
});

self.addEventListener('fetch', (e) => {
  e.respondWith(caches.match(e.request).then(res => res || fetch(e.request)));
});