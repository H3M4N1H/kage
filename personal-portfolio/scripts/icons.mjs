const paths = {
 arrow:'<path d="M7 17 17 7M7 7h10v10"/>',
 layers:'<path d="m3 7 9-5 9 5-9 5-9-5Zm0 5 9 5 9-5M3 17l9 5 9-5"/>',
 shield:'<path d="m12 3 8 3v6c0 5-8 9-8 9s-8-4-8-9V6l8-3Z"/><path d="m8 12 3 3 5-6"/>',
 globe:'<circle cx="12" cy="12" r="9"/><ellipse cx="12" cy="12" rx="4" ry="9"/><path d="M3 12h18M5 6h14M5 18h14"/>',
 network:'<circle cx="12" cy="4" r="2"/><circle cx="4" cy="19" r="2"/><circle cx="20" cy="19" r="2"/><path d="m11 6-6 11m8-11 6 11M6 19h12"/>',
 lock:'<rect x="5" y="10" width="14" height="12" rx="2"/><path d="M8 10V6a4 4 0 0 1 8 0v4m-4 5v3"/>',
 terminal:'<path d="m4 5 7 7-7 7m9 0h7"/>',
 code:'<path d="m7 6-6 6 6 6m10-12 6 6-6 6m-3-15-4 18"/>',
 flag:'<path d="M5 22V3m0 1c5-5 8 5 15 0v10c-7 5-10-5-15 0"/>',
 cpu:'<rect x="5" y="5" width="14" height="14" rx="2"/><path d="M9 1v4m6-4v4M9 19v4m6-4v4M1 9h4m-4 6h4m14-6h4m-4 6h4"/><rect x="9" y="9" width="6" height="6"/>',
 python:'<path d="M12 3H8a3 3 0 0 0-3 3v4h9v3H5a3 3 0 0 0-3 3v1a3 3 0 0 0 3 3h7m0-17h7a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-5v4a3 3 0 0 1-3 3H8"/>',
 linux:'<path d="M8 9V6a4 4 0 0 1 8 0v3l3 9-4 3h-6l-4-3 3-9Z"/><path d="m10 9 2 2 2-2M7 17l-3 4m13-4 3 4"/>',
 git:'<path d="m12 2 10 10-10 10L2 12 12 2Zm-4 4 8 8M9 7v10"/><circle cx="9" cy="17" r="1.5"/><circle cx="16" cy="14" r="1.5"/>',
 github:'<path d="M9 21c-5 1-5-3-7-3m14 5v-4c0-1-.3-2-1-2.5 4-.5 7-2 7-6a5 5 0 0 0-1.5-3.5c.2-1 .2-2-.5-3-2 0-3 1-4 1a13 13 0 0 0-8 0C7 4 6 3 4 3c-.7 1-.7 2-.5 3A5 5 0 0 0 2 9.5c0 4 3 5.5 7 6-.7.5-1 1.5-1 2.5v5"/>',
 linkedin:'<rect x="3" y="8" width="4" height="13"/><circle cx="5" cy="3" r="2"/><path d="M11 21V8h4v2c3-4 7-2 7 2v9h-4v-8c0-3-3-3-3 0v8Z"/>',
 cloud:'<path d="M7 18a5 5 0 1 1-1-10 7 7 0 0 1 13-1 5 5 0 0 1 0 10M8 21l2-3m3 4 2-3m3 3 2-3"/>',
 book:'<path d="M12 5C8 1 3 3 2 3v16c4-2 8 0 10 2 2-2 6-4 10-2V3c-1 0-6-2-10 2Zm0 0v16"/>',
 mail:'<rect x="2" y="4" width="20" height="16" rx="2"/><path d="m3 6 9 7 9-7"/>'
};
export const icon = (name, cls='') => `<svg class="icon ${cls}" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">${paths[name]||paths.code}</svg>`;
