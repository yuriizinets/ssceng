package action

var client = "<script>\"use strict\";(()=>{function C(e){let a=null;if(e.id&&(a=document.getElementById(e.id)),!a){let t=e.start,n=0;for(;t.parentElement;){if(!t.getAttribute(\"state\")){t=t.parentElement;continue}if(e.depth===n){a=t;break}t=t.parentElement,n++}}if(!a)throw new Error(`Root element not found with parameters ${e}.`);return a}var Y=11;function oe(e,a){var t=a.attributes,n,r,d,s,p;if(!(a.nodeType===Y||e.nodeType===Y)){for(var u=t.length-1;u>=0;u--)n=t[u],r=n.name,d=n.namespaceURI,s=n.value,d?(r=n.localName||r,p=e.getAttributeNS(d,r),p!==s&&(n.prefix===\"xmlns\"&&(r=n.name),e.setAttributeNS(d,r,s))):(p=e.getAttribute(r),p!==s&&e.setAttribute(r,s));for(var m=e.attributes,g=m.length-1;g>=0;g--)n=m[g],r=n.name,d=n.namespaceURI,d?(r=n.localName||r,a.hasAttributeNS(d,r)||e.removeAttributeNS(d,r)):a.hasAttribute(r)||e.removeAttribute(r)}}var B,se=\"http://www.w3.org/1999/xhtml\",h=typeof document>\"u\"?void 0:document,fe=!!h&&\"content\"in h.createElement(\"template\"),ue=!!h&&h.createRange&&\"createContextualFragment\"in h.createRange();function ce(e){var a=h.createElement(\"template\");return a.innerHTML=e,a.content.childNodes[0]}function pe(e){B||(B=h.createRange(),B.selectNode(h.body));var a=B.createContextualFragment(e);return a.childNodes[0]}function ve(e){var a=h.createElement(\"body\");return a.innerHTML=e,a.childNodes[0]}function he(e){return e=e.trim(),fe?ce(e):ue?pe(e):ve(e)}function P(e,a){var t=e.nodeName,n=a.nodeName,r,d;return t===n?!0:(r=t.charCodeAt(0),d=n.charCodeAt(0),r<=90&&d>=97?t===n.toUpperCase():d<=90&&r>=97?n===t.toUpperCase():!1)}function ge(e,a){return!a||a===se?h.createElement(e):h.createElementNS(a,e)}function me(e,a){for(var t=e.firstChild;t;){var n=t.nextSibling;a.appendChild(t),t=n}return a}function K(e,a,t){e[t]!==a[t]&&(e[t]=a[t],e[t]?e.setAttribute(t,\"\"):e.removeAttribute(t))}var Q={OPTION:function(e,a){var t=e.parentNode;if(t){var n=t.nodeName.toUpperCase();n===\"OPTGROUP\"&&(t=t.parentNode,n=t&&t.nodeName.toUpperCase()),n===\"SELECT\"&&!t.hasAttribute(\"multiple\")&&(e.hasAttribute(\"selected\")&&!a.selected&&(e.setAttribute(\"selected\",\"selected\"),e.removeAttribute(\"selected\")),t.selectedIndex=-1)}K(e,a,\"selected\")},INPUT:function(e,a){K(e,a,\"checked\"),K(e,a,\"disabled\"),e.value!==a.value&&(e.value=a.value),a.hasAttribute(\"value\")||e.removeAttribute(\"value\")},TEXTAREA:function(e,a){var t=a.value;e.value!==t&&(e.value=t);var n=e.firstChild;if(n){var r=n.nodeValue;if(r==t||!t&&r==e.placeholder)return;n.nodeValue=t}},SELECT:function(e,a){if(!a.hasAttribute(\"multiple\")){for(var t=-1,n=0,r=e.firstChild,d,s;r;)if(s=r.nodeName&&r.nodeName.toUpperCase(),s===\"OPTGROUP\")d=r,r=d.firstChild;else{if(s===\"OPTION\"){if(r.hasAttribute(\"selected\")){t=n;break}n++}r=r.nextSibling,!r&&d&&(r=d.nextSibling,d=null)}e.selectedIndex=t}}},L=1,Z=11,ee=3,te=8;function N(){}function Te(e){if(e)return e.getAttribute&&e.getAttribute(\"id\")||e.id}function Ae(e){return function(t,n,r){if(r||(r={}),typeof n==\"string\")if(t.nodeName===\"#document\"||t.nodeName===\"HTML\"||t.nodeName===\"BODY\"){var d=n;n=h.createElement(\"html\"),n.innerHTML=d}else n=he(n);else n.nodeType===Z&&(n=n.firstElementChild);var s=r.getNodeKey||Te,p=r.onBeforeNodeAdded||N,u=r.onNodeAdded||N,m=r.onBeforeElUpdated||N,g=r.onElUpdated||N,_=r.onBeforeNodeDiscarded||N,w=r.onNodeDiscarded||N,k=r.onBeforeElChildrenUpdated||N,ae=r.skipFromChildren||N,$=r.addChild||function(l,i){return l.appendChild(i)},V=r.childrenOnly===!0,y=Object.create(null),O=[];function x(l){O.push(l)}function z(l,i){if(l.nodeType===L)for(var c=l.firstChild;c;){var o=void 0;i&&(o=s(c))?x(o):(w(c),c.firstChild&&z(c,i)),c=c.nextSibling}}function H(l,i,c){_(l)!==!1&&(i&&i.removeChild(l),w(l),z(l,c))}function q(l){if(l.nodeType===L||l.nodeType===Z)for(var i=l.firstChild;i;){var c=s(i);c&&(y[c]=i),q(i),i=i.nextSibling}}q(t);function F(l){u(l);for(var i=l.firstChild;i;){var c=i.nextSibling,o=s(i);if(o){var f=y[o];f&&P(i,f)?(i.parentNode.replaceChild(f,i),U(f,i)):F(i)}else F(i);i=c}}function ie(l,i,c){for(;i;){var o=i.nextSibling;(c=s(i))?x(c):H(i,l,!0),i=o}}function U(l,i,c){var o=s(i);o&&delete y[o],!(!c&&(m(l,i)===!1||(e(l,i),g(l),k(l,i)===!1)))&&(l.nodeName!==\"TEXTAREA\"?le(l,i):Q.TEXTAREA(l,i))}function le(l,i){var c=ae(l),o=i.firstChild,f=l.firstChild,M,T,S,E,A;e:for(;o;){for(E=o.nextSibling,M=s(o);!c&&f;){if(S=f.nextSibling,o.isSameNode&&o.isSameNode(f)){o=E,f=S;continue e}T=s(f);var R=f.nodeType,b=void 0;if(R===o.nodeType&&(R===L?(M?M!==T&&((A=y[M])?S===A?b=!1:(l.insertBefore(A,f),T?x(T):H(f,l,!0),f=A):b=!1):T&&(b=!1),b=b!==!1&&P(f,o),b&&U(f,o)):(R===ee||R==te)&&(b=!0,f.nodeValue!==o.nodeValue&&(f.nodeValue=o.nodeValue))),b){o=E,f=S;continue e}T?x(T):H(f,l,!0),f=S}if(M&&(A=y[M])&&P(A,o))c||$(l,A),U(A,o);else{var G=p(o);G!==!1&&(G&&(o=G),o.actualize&&(o=o.actualize(l.ownerDocument||h)),$(l,o),F(o))}o=E,f=S}ie(l,f,T);var J=Q[l.nodeName];J&&J(l,i)}var v=t,D=v.nodeType,j=n.nodeType;if(!V){if(D===L)j===L?P(t,n)||(w(t),v=me(t,ge(n.nodeName,n.namespaceURI))):v=n;else if(D===ee||D===te){if(j===D)return v.nodeValue!==n.nodeValue&&(v.nodeValue=n.nodeValue),v;v=n}}if(v===n)w(t);else{if(n.isSameNode&&n.isSameNode(v))return;if(U(v,n,V),O)for(var I=0,de=O.length;I<de;I++){var W=y[O[I]];W&&H(W,W.parentNode,!1)}}return!V&&v!==t&&t.parentNode&&(v.actualize&&(v=v.actualize(t.ownerDocument||h)),t.parentNode.replaceChild(v,t)),v}}var be=Ae(oe),ne=be;function X(e,a,t){let n=new Array,r={};t&&(r=t),r.onBeforeElUpdated=function(d,s){if(d.getAttribute(\"action:morph.ignore.attr\")!=null){let p=d.getAttribute(\"ssa:morph.ignore.attr\");if(p)if(p==\"innerHTML\")s.innerHTML=d.innerHTML;else{let u=d.getAttribute(p);u&&s.setAttribute(p,u)}}return d.getAttribute(\"action:morph.ignore\")!=null?!1:d.getAttribute(\"action:morph.ignore.this\")!=null&&d!=e?(n.push({fromEl:d,toEl:s}),!1):!0},ne(e,a,r),n.length>0&&n.forEach(d=>{X(d.fromEl,d.toEl,{childrenOnly:!0})})}async function re(e,a,...t){let n=Ne(a),r=C({start:e,depth:n.depth,id:n.id});ye(r);let d=actionpath.endsWith(\"/\")?actionpath:actionpath+\"/\";d+=`${r.getAttribute(\"name\")}`,d+=`/${n.name}`;let s=new FormData;s.set(\"State\",r.getAttribute(\"state\")),s.set(\"Args\",JSON.stringify(t));let p=new XMLHttpRequest;p.open(\"POST\",d,!0);let u=\"\",m=0;p.onprogress=()=>{let g=p.responseText.length;if(m==g)return;let _=p.responseText.substring(m,g);if(u+=_,console.log(u,u.endsWith(actionterminator)),u.endsWith(actionterminator)){if(u=u.slice(0,-actionterminator.length),u.startsWith(\"action:redirect=\")){window.location.href=u.replace(\"action:redirect=\",\"\");return}switch(r.getAttribute(\"action:render.mode\")||\"morph\"){case\"replace\":r.outerHTML=u;break;case\"morph\":try{X(r,u)}catch(k){console.log('Fallback from \"morphdom\" to \"replace\" due to an error:',k),r.outerHTML=u}break;default:console.log('Render mode is not supported, fallback to \"replace\"'),r.outerHTML=u;break}u=\"\"}m=g},p.send(s)}function Ne(e){let a=e.split(\".\"),t=a[a.length-1],n=0,r;return e.startsWith(\"#\")&&(r=a[1]),e.startsWith(\"$\")&&(n=parseInt(a[1])),{name:t,depth:n,id:r}}function ye(e){let a=\"action\\\\:display\";e.querySelectorAll(`[${a}]`).forEach(n=>{let r=n.getAttribute(a);r&&(n.style.display=r)})}window.action=re;window.root=C;})();</script>"