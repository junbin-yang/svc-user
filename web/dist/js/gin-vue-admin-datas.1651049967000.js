/*! 
 Build Time : 1651049967000 */
import{s as a}from"./gin-vue-admin-authority.16510499670002.js";import{W as e}from"./gin-vue-admin-warningBar.1651049967000.js";import{r as t,b as l,o as r,c as u,d as o,e as i,w as s,G as n,H as d,i as h,j as c,l as y,t as m}from"../gva/gin-vue-admin-index.1651049967000.js";const v={class:"clearflex",style:{margin:"18px"}},f=h("确 定"),p=h("全选"),g=h("本角色"),I=h("本角色及子角色"),w={name:"Datas"},x=Object.assign(w,{props:{row:{default:function(){return{}},type:Object},authority:{default:function(){return[]},type:Array}},emits:["changeRow"],setup(w,{expose:x,emit:b}){const A=w,k=t([]),j=t(!1),C=a=>{a&&a.forEach((a=>{const e={};e.authorityId=a.authorityId,e.authorityName=a.authorityName,k.value.push(e),a.children&&a.children.length&&C(a.children)}))},_=t([]);C(A.authority),A.row.dataAuthorityId&&A.row.dataAuthorityId.forEach((a=>{const e=k.value&&k.value.filter((e=>e.authorityId===a.authorityId))&&k.value.filter((e=>e.authorityId===a.authorityId))[0];_.value.push(e)}));const R=()=>{_.value=[...k.value],b("changeRow","dataAuthorityId",_.value),j.value=!0},z=()=>{_.value=k.value.filter((a=>a.authorityId===A.row.authorityId)),b("changeRow","dataAuthorityId",_.value),j.value=!0},N=()=>{const a=[];E(A.row,a),_.value=k.value.filter((e=>a.indexOf(e.authorityId)>-1)),b("changeRow","dataAuthorityId",_.value),j.value=!0},E=(a,e)=>{e.push(a.authorityId),a.children&&a.children.forEach((a=>{E(a,e)}))},O=()=>{return e=this,t=null,l=function*(){0===(yield a(A.row)).code&&c({type:"success",message:"资源设置成功"})},new Promise(((a,r)=>{var u=a=>{try{i(l.next(a))}catch(e){r(e)}},o=a=>{try{i(l.throw(a))}catch(e){r(e)}},i=e=>e.done?a(e.value):Promise.resolve(e.value).then(u,o);i((l=l.apply(e,t)).next())}));var e,t,l},V=()=>{b("changeRow","dataAuthorityId",_.value),j.value=!0};return x({enterAndNext:()=>{O()},needConfirm:j}),(a,t)=>{const c=l("el-button"),w=l("el-checkbox"),x=l("el-checkbox-group");return r(),u("div",null,[o("div",v,[i(c,{class:"fl-right",size:"small",type:"primary",onClick:O},{default:s((()=>[f])),_:1}),i(c,{class:"fl-left",size:"small",type:"primary",onClick:R},{default:s((()=>[p])),_:1}),i(c,{class:"fl-left",size:"small",type:"primary",onClick:z},{default:s((()=>[g])),_:1}),i(c,{class:"fl-left",size:"small",type:"primary",onClick:N},{default:s((()=>[I])),_:1})]),i(x,{modelValue:_.value,"onUpdate:modelValue":t[0]||(t[0]=a=>_.value=a),onChange:V},{default:s((()=>[(r(!0),u(n,null,d(k.value,((a,e)=>(r(),y(w,{key:e,label:a},{default:s((()=>[h(m(a.authorityName),1)])),_:2},1032,["label"])))),128))])),_:1},8,["modelValue"]),i(e,{title:"此功能仅用于创建角色和角色的many2many关系表，具体使用还须自己结合表实现业务，详情参考示例代码（客户示例）"})])}}});export{x as default};