/*! 
 Build Time : 1650194964000 */
import{C as a}from"./gin-vue-admin-github.1650194964000.js";import{f as s}from"./gin-vue-admin-date.1650194964000.js";/* empty css                                                                                      */import{_ as e,r as t,o as m,c as o,d as i,G as l,H as d,p as n,q as c,n as r,t as v}from"../gva/gin-vue-admin-index.1650194964000.js";const u={class:"commit-table"},f=(a=>(n("data-v-02421634"),a=a(),c(),a))((()=>i("div",{class:"commit-table-title"}," 更新日志 ",-1))),g={class:"log"},p={class:"flex-1 flex key-box"},_={class:"flex-5 flex message"},h={class:"flex-3 flex form"},x={name:"DashboardTable"};var y=e(Object.assign(x,{setup(e){const n=t(!0),c=t([]);return a(0).then((({data:a})=>{n.value=!1,a.forEach(((a,e)=>{a.commit.message&&e<10&&c.value.push({from:s(a.commit.author.date,"yyyy-MM-dd"),title:a.commit.author.name,showDayAndMonth:!0,message:a.commit.message})}))})),(a,s)=>(m(),o("div",u,[f,i("div",g,[(m(!0),o(l,null,d(c.value,((a,s)=>(m(),o("div",{key:s,class:"log-item"},[i("div",p,[i("span",{class:r(["key",s<3&&"top"])},v(s+1),3)]),i("div",_,v(a.message),1),i("div",h,v(a.from),1)])))),128))])]))}}),[["__scopeId","data-v-02421634"]]);export{y as default};
