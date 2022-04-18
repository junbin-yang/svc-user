/*! 
 Build Time : 1650194964000 */
import{s as e,_ as a,u as l,r as s,a as t,b as o,o as d,c as i,d as p,n as u,e as n,w as r,f as m,g as c,h as b,i as v,E as h,j as g}from"../gva/gin-vue-admin-index.1650194964000.js";const y={class:"init_page"},f={class:"init_page_panle"},_=b('<div class="hello_title" data-v-1bb614c4>GIN-VUE-ADMIN</div><p class="in-two a-fadeinT" data-v-1bb614c4>初始化须知</p><p class="init_p" data-v-1bb614c4>1.您需有用一定的VUE和GOLANG基础</p><p class="init_p" data-v-1bb614c4>2.请您确认是否已经阅读过官方文档</p><p class="init_p" data-v-1bb614c4>3.请您确认是否了解后续的配置流程</p><p class="init_p" data-v-1bb614c4>注：开发组不为文档中书写过的内容提供无偿服务</p>',6),V={class:"init_btn"},w=v(" 阅读文档 "),N=v(" 我已确认 "),q={style:{"text-align":"right"}},k=v("立即初始化"),U={name:"Init"};var x=a(Object.assign(U,{setup(a){const b=l(),v=s(0),U=()=>{v.value=v.value+1},x=()=>{window.open("https://www.gin-vue-admin.com/docs/first_master#3-init")},T=s(!1),j=t({dbType:"mysql",host:"127.0.0.1",port:"3306",userName:"root",password:"",dbName:"gva"}),C=e=>{switch(e){case"mysql":default:Object.assign(j,{dbType:"mysql",host:"127.0.0.1",port:"3306",userName:"root",password:"",dbName:"gva"});break;case"pgsql":Object.assign(j,{dbType:"pgsql",host:"127.0.0.1",port:"5432",userName:"postgres",password:"",dbName:"gva"})}},I=()=>{return a=this,l=null,s=function*(){const a=h.service({lock:!0,text:"正在初始化数据库，请稍候",spinner:"loading",background:"rgba(0, 0, 0, 0.7)"});try{const s=yield(l=j,e({url:"/init/initdb",method:"post",data:l}));0===s.code&&(T.value=!0,g({type:"success",message:s.msg}),b.push({name:"Login"})),a.close()}catch(s){a.close()}var l},new Promise(((e,t)=>{var o=e=>{try{i(s.next(e))}catch(a){t(a)}},d=e=>{try{i(s.throw(e))}catch(a){t(a)}},i=a=>a.done?e(a.value):Promise.resolve(a.value).then(o,d);i((s=s.apply(a,l)).next())}));var a,l,s};return(e,a)=>{const l=o("el-button"),s=o("el-option"),t=o("el-select"),b=o("el-form-item"),h=o("el-input"),g=o("el-form");return d(),i("div",y,[p("div",f,[v.value<2?(d(),i("div",{key:0,id:"hello",class:u([[v.value<1?"slide-in-fwd-top":"slide-out-right"],"hello"])},[p("div",null,[_,p("p",V,[n(l,{type:"primary",onClick:x},{default:r((()=>[w])),_:1}),n(l,{type:"primary",onClick:U},{default:r((()=>[N])),_:1})])])],2)):m("",!0),v.value>0?(d(),i("div",{key:1,class:u([[v.value>0&&!T.value?"slide-in-left":"",T.value?"slide-out-right":""],"form"])},[n(g,{ref:"formRef",model:c(j),"label-width":"100px"},{default:r((()=>[n(b,{label:"数据库类型"},{default:r((()=>[n(t,{modelValue:c(j).dbType,"onUpdate:modelValue":a[0]||(a[0]=e=>c(j).dbType=e),placeholder:"请选择",onChange:C},{default:r((()=>[n(s,{key:"mysql",label:"mysql",value:"mysql"}),n(s,{key:"pgsql",label:"pgsql(测试版)",value:"pgsql"})])),_:1},8,["modelValue"])])),_:1}),n(b,{label:"host"},{default:r((()=>[n(h,{modelValue:c(j).host,"onUpdate:modelValue":a[1]||(a[1]=e=>c(j).host=e),placeholder:"请输入数据库链接"},null,8,["modelValue"])])),_:1}),n(b,{label:"port"},{default:r((()=>[n(h,{modelValue:c(j).port,"onUpdate:modelValue":a[2]||(a[2]=e=>c(j).port=e),placeholder:"请输入数据库端口"},null,8,["modelValue"])])),_:1}),n(b,{label:"userName"},{default:r((()=>[n(h,{modelValue:c(j).userName,"onUpdate:modelValue":a[3]||(a[3]=e=>c(j).userName=e),placeholder:"请输入数据库用户名"},null,8,["modelValue"])])),_:1}),n(b,{label:"password"},{default:r((()=>[n(h,{modelValue:c(j).password,"onUpdate:modelValue":a[4]||(a[4]=e=>c(j).password=e),placeholder:"请输入数据库密码（没有则为空）"},null,8,["modelValue"])])),_:1}),n(b,{label:"dbName"},{default:r((()=>[n(h,{modelValue:c(j).dbName,"onUpdate:modelValue":a[5]||(a[5]=e=>c(j).dbName=e),placeholder:"请输入数据库名称"},null,8,["modelValue"])])),_:1}),n(b,null,{default:r((()=>[p("div",q,[n(l,{type:"primary",onClick:I},{default:r((()=>[k])),_:1})])])),_:1})])),_:1},8,["model"])],2)):m("",!0)])])}}}),[["__scopeId","data-v-1bb614c4"]]);export{x as default};
