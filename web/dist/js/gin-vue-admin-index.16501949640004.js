/*! 
 Build Time : 1650194964000 */
var t=(t,e,n)=>new Promise(((r,o)=>{var a=t=>{try{c(n.next(t))}catch(e){o(e)}},i=t=>{try{c(n.throw(t))}catch(e){o(e)}},c=t=>t.done?r(t.value):Promise.resolve(t.value).then(a,i);c((n=n.apply(t,e)).next())}));import"./gin-vue-admin-echartsLine.vue_vue_type_style_index_0_scoped_true_lang.1650194964000.js";import{D as e,N as n,_ as r,r as o,u as a,b as i,o as c,c as u,d as l,e as s,w as d,i as f,t as p,G as y,H as v,f as m,p as g,q as h,O as b,P as x,j as _,Q as w,l as E,R as S,x as T}from"../gva/gin-vue-admin-index.1650194964000.js";import"./gin-vue-admin-github.1650194964000.js";import"./gin-vue-admin-date.1650194964000.js";/* empty css                                                                                      */var k,j={exports:{}};k=function(){return function(){var t={686:function(t,e,n){n.d(e,{default:function(){return E}});var r=n(279),o=n.n(r),a=n(370),i=n.n(a),c=n(817),u=n.n(c);function l(t){try{return document.execCommand(t)}catch(e){return!1}}var s=function(t){var e=u()(t);return l("cut"),e};function d(t){var e="rtl"===document.documentElement.getAttribute("dir"),n=document.createElement("textarea");n.style.fontSize="12pt",n.style.border="0",n.style.padding="0",n.style.margin="0",n.style.position="absolute",n.style[e?"right":"left"]="-9999px";var r=window.pageYOffset||document.documentElement.scrollTop;return n.style.top="".concat(r,"px"),n.setAttribute("readonly",""),n.value=t,n}var f=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{container:document.body},n="";if("string"==typeof t){var r=d(t);e.container.appendChild(r),n=u()(r),l("copy"),r.remove()}else n=u()(t),l("copy");return n};function p(t){return(p="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t})(t)}var y=function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},e=t.action,n=void 0===e?"copy":e,r=t.container,o=t.target,a=t.text;if("copy"!==n&&"cut"!==n)throw new Error('Invalid "action" value, use either "copy" or "cut"');if(void 0!==o){if(!o||"object"!==p(o)||1!==o.nodeType)throw new Error('Invalid "target" value, use a valid Element');if("copy"===n&&o.hasAttribute("disabled"))throw new Error('Invalid "target" attribute. Please use "readonly" instead of "disabled" attribute');if("cut"===n&&(o.hasAttribute("readonly")||o.hasAttribute("disabled")))throw new Error('Invalid "target" attribute. You can\'t cut text from elements with "readonly" or "disabled" attributes')}return a?f(a,{container:r}):o?"cut"===n?s(o):f(o,{container:r}):void 0};function v(t){return(v="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t})(t)}function m(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(t,r.key,r)}}function g(t,e){return(g=Object.setPrototypeOf||function(t,e){return t.__proto__=e,t})(t,e)}function h(t){var e=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(t){return!1}}();return function(){var n,r=x(t);if(e){var o=x(this).constructor;n=Reflect.construct(r,arguments,o)}else n=r.apply(this,arguments);return b(this,n)}}function b(t,e){return!e||"object"!==v(e)&&"function"!=typeof e?function(t){if(void 0===t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return t}(t):e}function x(t){return(x=Object.setPrototypeOf?Object.getPrototypeOf:function(t){return t.__proto__||Object.getPrototypeOf(t)})(t)}function _(t,e){var n="data-clipboard-".concat(t);if(e.hasAttribute(n))return e.getAttribute(n)}var w=function(t){!function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function");t.prototype=Object.create(e&&e.prototype,{constructor:{value:t,writable:!0,configurable:!0}}),e&&g(t,e)}(a,t);var e,n,r,o=h(a);function a(t,e){var n;return function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}(this,a),(n=o.call(this)).resolveOptions(e),n.listenClick(t),n}return e=a,n=[{key:"resolveOptions",value:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};this.action="function"==typeof t.action?t.action:this.defaultAction,this.target="function"==typeof t.target?t.target:this.defaultTarget,this.text="function"==typeof t.text?t.text:this.defaultText,this.container="object"===v(t.container)?t.container:document.body}},{key:"listenClick",value:function(t){var e=this;this.listener=i()(t,"click",(function(t){return e.onClick(t)}))}},{key:"onClick",value:function(t){var e=t.delegateTarget||t.currentTarget,n=this.action(e)||"copy",r=y({action:n,container:this.container,target:this.target(e),text:this.text(e)});this.emit(r?"success":"error",{action:n,text:r,trigger:e,clearSelection:function(){e&&e.focus(),document.activeElement.blur(),window.getSelection().removeAllRanges()}})}},{key:"defaultAction",value:function(t){return _("action",t)}},{key:"defaultTarget",value:function(t){var e=_("target",t);if(e)return document.querySelector(e)}},{key:"defaultText",value:function(t){return _("text",t)}},{key:"destroy",value:function(){this.listener.destroy()}}],r=[{key:"copy",value:function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{container:document.body};return f(t,e)}},{key:"cut",value:function(t){return s(t)}},{key:"isSupported",value:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:["copy","cut"],e="string"==typeof t?[t]:t,n=!!document.queryCommandSupported;return e.forEach((function(t){n=n&&!!document.queryCommandSupported(t)})),n}}],n&&m(e.prototype,n),r&&m(e,r),a}(o()),E=w},828:function(t){if("undefined"!=typeof Element&&!Element.prototype.matches){var e=Element.prototype;e.matches=e.matchesSelector||e.mozMatchesSelector||e.msMatchesSelector||e.oMatchesSelector||e.webkitMatchesSelector}t.exports=function(t,e){for(;t&&9!==t.nodeType;){if("function"==typeof t.matches&&t.matches(e))return t;t=t.parentNode}}},438:function(t,e,n){var r=n(828);function o(t,e,n,r,o){var i=a.apply(this,arguments);return t.addEventListener(n,i,o),{destroy:function(){t.removeEventListener(n,i,o)}}}function a(t,e,n,o){return function(n){n.delegateTarget=r(n.target,e),n.delegateTarget&&o.call(t,n)}}t.exports=function(t,e,n,r,a){return"function"==typeof t.addEventListener?o.apply(null,arguments):"function"==typeof n?o.bind(null,document).apply(null,arguments):("string"==typeof t&&(t=document.querySelectorAll(t)),Array.prototype.map.call(t,(function(t){return o(t,e,n,r,a)})))}},879:function(t,e){e.node=function(t){return void 0!==t&&t instanceof HTMLElement&&1===t.nodeType},e.nodeList=function(t){var n=Object.prototype.toString.call(t);return void 0!==t&&("[object NodeList]"===n||"[object HTMLCollection]"===n)&&"length"in t&&(0===t.length||e.node(t[0]))},e.string=function(t){return"string"==typeof t||t instanceof String},e.fn=function(t){return"[object Function]"===Object.prototype.toString.call(t)}},370:function(t,e,n){var r=n(879),o=n(438);t.exports=function(t,e,n){if(!t&&!e&&!n)throw new Error("Missing required arguments");if(!r.string(e))throw new TypeError("Second argument must be a String");if(!r.fn(n))throw new TypeError("Third argument must be a Function");if(r.node(t))return function(t,e,n){return t.addEventListener(e,n),{destroy:function(){t.removeEventListener(e,n)}}}(t,e,n);if(r.nodeList(t))return function(t,e,n){return Array.prototype.forEach.call(t,(function(t){t.addEventListener(e,n)})),{destroy:function(){Array.prototype.forEach.call(t,(function(t){t.removeEventListener(e,n)}))}}}(t,e,n);if(r.string(t))return function(t,e,n){return o(document.body,t,e,n)}(t,e,n);throw new TypeError("First argument must be a String, HTMLElement, HTMLCollection, or NodeList")}},817:function(t){t.exports=function(t){var e;if("SELECT"===t.nodeName)t.focus(),e=t.value;else if("INPUT"===t.nodeName||"TEXTAREA"===t.nodeName){var n=t.hasAttribute("readonly");n||t.setAttribute("readonly",""),t.select(),t.setSelectionRange(0,t.value.length),n||t.removeAttribute("readonly"),e=t.value}else{t.hasAttribute("contenteditable")&&t.focus();var r=window.getSelection(),o=document.createRange();o.selectNodeContents(t),r.removeAllRanges(),r.addRange(o),e=r.toString()}return e}},279:function(t){function e(){}e.prototype={on:function(t,e,n){var r=this.e||(this.e={});return(r[t]||(r[t]=[])).push({fn:e,ctx:n}),this},once:function(t,e,n){var r=this;function o(){r.off(t,o),e.apply(n,arguments)}return o._=e,this.on(t,o,n)},emit:function(t){for(var e=[].slice.call(arguments,1),n=((this.e||(this.e={}))[t]||[]).slice(),r=0,o=n.length;r<o;r++)n[r].fn.apply(n[r].ctx,e);return this},off:function(t,e){var n=this.e||(this.e={}),r=n[t],o=[];if(r&&e)for(var a=0,i=r.length;a<i;a++)r[a].fn!==e&&r[a].fn._!==e&&o.push(r[a]);return o.length?n[t]=o:delete n[t],this}},t.exports=e,t.exports.TinyEmitter=e}},e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={exports:{}};return t[r](o,o.exports,n),o.exports}return n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,{a:e}),e},n.d=function(t,e){for(var r in e)n.o(e,r)&&!n.o(t,r)&&Object.defineProperty(t,r,{enumerable:!0,get:e[r]})},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n(686)}().default};var A=n(j.exports=k());const C=t=>(g("data-v-cf7fdd26"),t=t(),h(),t),O={class:"page"},L={class:"gva-card-box"},P={class:"gva-card gva-top-card"},R={class:"gva-top-card-left"},q=C((()=>l("div",{class:"gva-top-card-left-title"},"早安，管理员，请开始一天的工作吧",-1))),M=C((()=>l("div",{class:"gva-top-card-left-dot"},[l("br")],-1))),N={class:"gva-top-card-left-rows"},I={class:"flex-center"},H={class:"flex-center"},D={class:"gva-top-card-left-item"},F=f(" 用户注册： "),z=f("生成邀请码"),B=C((()=>l("img",{src:"./assets/dashboard.70e55b71.png",class:"gva-top-card-right",alt:""},null,-1))),Y={class:"gva-card-box"},$=C((()=>l("div",{class:"card-header"},[l("span",null,"快捷入口")],-1))),G={class:"quick-entrance-item"},Q={name:"Dashboard"};var U=r(Object.assign(Q,{setup(e){const{toClipboard:n}=(t=>{const e=void 0===(null==t?void 0:t.appendToBody)||t.appendToBody;return{toClipboard:(t,n)=>new Promise(((r,o)=>{const a=document.createElement("button"),i=new A(a,{text:()=>t,action:()=>"copy",container:void 0!==n?n:document.body});i.on("success",(t=>{i.destroy(),r(t)})),i.on("error",(t=>{i.destroy(),o(t)})),e&&document.body.appendChild(a),a.click(),e&&document.body.removeChild(a)}))}})(),r=()=>t(this,null,(function*(){const t=yield b();200===t.code?(n(t.data),x.alert(t.data,"邀请码")):_({type:"error",message:t.message})})),g=o({userTotal:0,loginTaday:0});(()=>{t(this,null,(function*(){const t=yield w();200===t.code&&(g.value=t.data)}))})();const h=o([{label:"应用管理",icon:"monitor",name:"applist",color:"#ff9c6e",bg:"rgba(255, 156, 110,.3)"},{label:"用户管理",icon:"user",name:"userlist",color:"#5cdbd3",bg:"rgba(92, 219, 211,.3)"}]),k=a();return(t,e)=>{const n=i("sort"),o=i("el-icon"),a=i("el-col"),b=i("avatar"),x=i("el-row"),_=i("el-button"),w=i("el-card");return c(),u("div",O,[l("div",L,[l("div",P,[l("div",R,[q,M,l("div",N,[s(x,null,{default:d((()=>[s(a,{span:8,xs:24,sm:8},{default:d((()=>[l("div",I,[s(o,{class:"dasboard-icon"},{default:d((()=>[s(n)])),_:1}),f(" "+p(`今日登录 (${g.value.loginTaday})`),1)])])),_:1}),s(a,{span:8,xs:24,sm:8},{default:d((()=>[l("div",H,[s(o,{class:"dasboard-icon"},{default:d((()=>[s(b)])),_:1}),f(" "+p(`总用户数 (${g.value.userTotal})`),1)])])),_:1})])),_:1})]),l("div",null,[l("div",D,[F,s(_,{color:"#626aef",plain:"",onClick:r},{default:d((()=>[z])),_:1})])])]),B])]),l("div",Y,[s(w,{class:"gva-card quick-entrance"},{header:d((()=>[$])),default:d((()=>[s(x,{gutter:20},{default:d((()=>[(c(!0),u(y,null,v(h.value,((t,e)=>(c(),E(a,{key:e,span:4,xs:8,class:"quick-entrance-items",onClick:e=>{return n=t.name,void k.push({name:n});var n}},{default:d((()=>[l("div",G,[l("div",{class:"quick-entrance-item-icon",style:S({backgroundColor:t.bg})},[s(o,null,{default:d((()=>[(c(),E(T(t.icon),{style:S({color:t.color})},null,8,["style"]))])),_:2},1024)],4),l("p",null,p(t.label),1)])])),_:2},1032,["onClick"])))),128))])),_:1})])),_:1})]),m("",!0)])}}}),[["__scopeId","data-v-cf7fdd26"]]);export{U as default};
