
const code = "body{margin:0;font-family:Arial, Helvetica, sans-serif;background-color:#F9F6F6}.App.svelte-1efkygv{text-align:center}.App-header.svelte-1efkygv{color:#333;display:flex;flex-direction:column;align-items:center;justify-content:center;font-size:calc(10px + 2vmin)}";

const styleEl = document.createElement("style");
const codeEl = document.createTextNode(code);
styleEl.type = 'text/css';

styleEl.appendChild(codeEl);
document.head.appendChild(styleEl);
