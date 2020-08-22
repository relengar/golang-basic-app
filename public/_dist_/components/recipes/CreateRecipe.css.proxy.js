
const code = ".button.svelte-1b2rh29{width:150px;border:1px solid rgba(129, 88, 12, 0.37);border-radius:10px;font-weight:500;font-size:18px;padding:3px 0 3px 0}";

const styleEl = document.createElement("style");
const codeEl = document.createTextNode(code);
styleEl.type = 'text/css';

styleEl.appendChild(codeEl);
document.head.appendChild(styleEl);
