/*! 
 Build Time : 1651049967000 */
import e from"./gin-vue-admin-menuItem.1651049967000.js";import n from"./gin-vue-admin-asyncSubmenu.1651049967000.js";import{V as r,b as o,o as t,l as u,w as i,c as s,H as a,G as f,f as d,x as l,g as m}from"../gva/gin-vue-admin-index.1651049967000.js";const c={name:"AsideComponent"},p=Object.assign(c,{props:{routerInfo:{type:Object,default:()=>null}},setup(c){const p=c,h=r((()=>p.routerInfo.children&&p.routerInfo.children.filter((e=>!e.hidden)).length?n:e));return(e,n)=>{const r=o("AsideComponent");return c.routerInfo.hidden?d("",!0):(t(),u(l(m(h)),{key:0,"router-info":c.routerInfo},{default:i((()=>[c.routerInfo.children&&c.routerInfo.children.length?(t(!0),s(f,{key:0},a(c.routerInfo.children,(e=>(t(),u(r,{key:e.name,"router-info":e},null,8,["router-info"])))),128)):d("",!0)])),_:1},8,["router-info"]))}}});export{p as default};