/*! 
 Build Time : 1651049967000 */
System.register(["./gin-vue-admin-github-legacy.1651049967000.js","./gin-vue-admin-date-legacy.1651049967000.js","./gin-vue-admin-dashboardTable.vue_vue_type_style_index_0_scoped_true_lang-legacy.1651049967000.js","../gva/gin-vue-admin-index-legacy.1651049967000.js"],(function(e){"use strict";var a,t,n,s,i,c,u,o,l,r,d,m,f;return{setters:[function(e){a=e.C},function(e){t=e.f},function(){},function(e){n=e._,s=e.r,i=e.o,c=e.c,u=e.d,o=e.G,l=e.H,r=e.p,d=e.q,m=e.n,f=e.t}],execute:function(){var v={class:"commit-table"},g=function(e){return r("data-v-6c1afbee"),e=e(),d(),e}((function(){return u("div",{class:"commit-table-title"}," 更新日志 ",-1)})),y={class:"log"},b={class:"flex-1 flex key-box"},_={class:"flex-5 flex message"},h={class:"flex-3 flex form"},x={name:"DashboardTable"};e("default",n(Object.assign(x,{setup:function(e){var n=s(!0),r=s([]);return a(0).then((function(e){var a=e.data;n.value=!1,a.forEach((function(e,a){e.commit.message&&a<10&&r.value.push({from:t(e.commit.author.date,"yyyy-MM-dd"),title:e.commit.author.name,showDayAndMonth:!0,message:e.commit.message})}))})),function(e,a){return i(),c("div",v,[g,u("div",y,[(i(!0),c(o,null,l(r.value,(function(e,a){return i(),c("div",{key:a,class:"log-item"},[u("div",b,[u("span",{class:m(["key",a<3&&"top"])},f(a+1),3)]),u("div",_,f(e.message),1),u("div",h,f(e.from),1)])})),128))])])}}}),[["__scopeId","data-v-6c1afbee"]]))}}}));