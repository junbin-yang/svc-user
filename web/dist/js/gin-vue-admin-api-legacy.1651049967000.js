/*! 
 Build Time : 1651049967000 */
!function(){function e(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function t(t){for(var r=1;r<arguments.length;r++){var a=null!=arguments[r]?arguments[r]:{};r%2?e(Object(a),!0).forEach((function(e){n(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):e(Object(a)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}function n(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t,n,r,a,u,l){try{var o=e[u](l),i=o.value}catch(c){return void n(c)}o.done?t(i):Promise.resolve(i).then(r,a)}function a(e){return function(){var t=this,n=arguments;return new Promise((function(a,u){var l=e.apply(t,n);function o(e){r(l,a,u,o,i,"next",e)}function i(e){r(l,a,u,o,i,"throw",e)}o(void 0)}))}}var u=document.createElement("style");u.innerHTML=".button-box[data-v-7b528490]{padding:10px 20px}.button-box .el-button[data-v-7b528490]{float:right}.warning[data-v-7b528490]{color:#dc143c}\n",document.head.appendChild(u),System.register(["./gin-vue-admin-api-legacy.16510499670002.js","./gin-vue-admin-stringFun-legacy.1651049967000.js","./gin-vue-admin-warningBar-legacy.1651049967000.js","../gva/gin-vue-admin-index-legacy.1651049967000.js"],(function(e){"use strict";var n,r,u,l,o,i,c,s,p,d,f,v,m,g,b,h,y,w,x,_,k,V,C,O,j;return{setters:[function(e){n=e.g,r=e.d,u=e.a,l=e.u,o=e.c,i=e.b},function(e){c=e.t},function(e){s=e.W},function(e){p=e._,d=e.r,f=e.b,v=e.o,m=e.c,g=e.d,b=e.e,h=e.w,y=e.G,w=e.H,x=e.t,_=e.i,k=e.p,V=e.q,C=e.j,O=e.P,j=e.l}],execute:function(){var P=function(e){return k("data-v-7b528490"),e=e(),V(),e},z={class:"gva-search-box"},R=_("查询"),A=_("重置"),G={class:"gva-table-box"},U={class:"gva-btn-list"},I=_("新增"),D=P((function(){return g("p",null,"确定要删除吗？",-1)})),S={style:{"text-align":"right","margin-top":"8px"}},E=_("取消"),T=_("确定"),q=_("删除"),F=_("编辑"),B=_("删除"),H={class:"gva-pagination"},L={class:"dialog-footer"},K=_("取 消"),M=_("确 定"),W={name:"Api"},J=Object.assign(W,{setup:function(e){var p=d([]),_=d({path:"",apiGroup:"",method:"",description:""}),k=d([{value:"POST",label:"创建",type:"success"},{value:"GET",label:"查看",type:""},{value:"PUT",label:"更新",type:"warning"},{value:"DELETE",label:"删除",type:"danger"}]),V=d(""),P=d({path:[{required:!0,message:"请输入api路径",trigger:"blur"}],apiGroup:[{required:!0,message:"请输入组名称",trigger:"blur"}],method:[{required:!0,message:"请选择请求方式",trigger:"blur"}],description:[{required:!0,message:"请输入api介绍",trigger:"blur"}]}),W=d(1),J=d(0),N=d(10),Q=d([]),X=d({}),Y=function(){X.value={}},Z=function(){W.value=1,N.value=10,ne()},$=function(e){N.value=e,ne()},ee=function(e){W.value=e,ne()},te=function(e){var t=e.prop,n=e.order;t&&("ID"===t&&(t="id"),X.value.orderKey=c(t),X.value.desc="descending"===n),ne()},ne=function(){var e=a(regeneratorRuntime.mark((function e(){var r;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,n(t({page:W.value,pageSize:N.value},X.value));case 2:0===(r=e.sent).code&&(Q.value=r.data.list,J.value=r.data.total,W.value=r.data.page,N.value=r.data.pageSize);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();ne();var re=function(e){p.value=e},ae=d(!1),ue=function(){var e=a(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=p.value.map((function(e){return e.ID})),e.next=3,r({ids:t});case 3:0===(n=e.sent).code&&(C({type:"success",message:n.msg}),Q.value.length===t.length&&W.value>1&&W.value--,ae.value=!1,ne());case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),le=d(null),oe=d("新增Api"),ie=d(!1),ce=function(e){switch(e){case"addApi":oe.value="新增Api";break;case"edit":oe.value="编辑Api"}V.value=e,ie.value=!0},se=function(){le.value.resetFields(),_.value={path:"",apiGroup:"",method:"",description:""},ie.value=!1},pe=function(){var e=a(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,u({id:t.ID});case 2:n=e.sent,_.value=n.data.api,ce("edit");case 5:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),de=function(){var e=a(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:le.value.validate(function(){var e=a(regeneratorRuntime.mark((function e(t){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!t){e.next=20;break}e.t0=V.value,e.next="addApi"===e.t0?4:"edit"===e.t0?11:18;break;case 4:return e.next=6,o(_.value);case 6:return 0===e.sent.code&&C({type:"success",message:"添加成功",showClose:!0}),ne(),se(),e.abrupt("break",20);case 11:return e.next=13,l(_.value);case 13:return 0===e.sent.code&&C({type:"success",message:"编辑成功",showClose:!0}),ne(),se(),e.abrupt("break",20);case 18:return C({type:"error",message:"未知操作",showClose:!0}),e.abrupt("break",20);case 20:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}());case 1:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),fe=function(){var e=a(regeneratorRuntime.mark((function e(t){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:O.confirm("此操作将永久删除所有角色下该api, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(a(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,i(t);case 2:0===e.sent.code&&(C({type:"success",message:"删除成功!"}),1===Q.value.length&&W.value>1&&W.value--,ne());case 4:case"end":return e.stop()}}),e)}))));case 1:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}();return function(e,t){var n=f("el-input"),r=f("el-form-item"),a=f("el-option"),u=f("el-select"),l=f("el-button"),o=f("el-form"),i=f("el-popover"),c=f("el-table-column"),d=f("el-table"),V=f("el-pagination"),C=f("el-dialog");return v(),m("div",null,[g("div",z,[b(o,{ref:"searchForm",inline:!0,model:X.value},{default:h((function(){return[b(r,{label:"路径"},{default:h((function(){return[b(n,{modelValue:X.value.path,"onUpdate:modelValue":t[0]||(t[0]=function(e){return X.value.path=e}),placeholder:"路径"},null,8,["modelValue"])]})),_:1}),b(r,{label:"描述"},{default:h((function(){return[b(n,{modelValue:X.value.description,"onUpdate:modelValue":t[1]||(t[1]=function(e){return X.value.description=e}),placeholder:"描述"},null,8,["modelValue"])]})),_:1}),b(r,{label:"API组"},{default:h((function(){return[b(n,{modelValue:X.value.apiGroup,"onUpdate:modelValue":t[2]||(t[2]=function(e){return X.value.apiGroup=e}),placeholder:"api组"},null,8,["modelValue"])]})),_:1}),b(r,{label:"请求"},{default:h((function(){return[b(u,{modelValue:X.value.method,"onUpdate:modelValue":t[3]||(t[3]=function(e){return X.value.method=e}),clearable:"",placeholder:"请选择"},{default:h((function(){return[(v(!0),m(y,null,w(k.value,(function(e){return v(),j(a,{key:e.value,label:"".concat(e.label,"(").concat(e.value,")"),value:e.value},null,8,["label","value"])})),128))]})),_:1},8,["modelValue"])]})),_:1}),b(r,null,{default:h((function(){return[b(l,{size:"small",type:"primary",icon:"search",onClick:Z},{default:h((function(){return[R]})),_:1}),b(l,{size:"small",icon:"refresh",onClick:Y},{default:h((function(){return[A]})),_:1})]})),_:1})]})),_:1},8,["model"])]),g("div",G,[g("div",U,[b(l,{size:"small",type:"primary",icon:"plus",onClick:t[4]||(t[4]=function(e){return ce("addApi")})},{default:h((function(){return[I]})),_:1}),b(i,{visible:ae.value,"onUpdate:visible":t[7]||(t[7]=function(e){return ae.value=e}),placement:"top",width:"160"},{reference:h((function(){return[b(l,{icon:"delete",size:"small",disabled:!p.value.length,style:{"margin-left":"10px"},onClick:t[6]||(t[6]=function(e){return ae.value=!0})},{default:h((function(){return[q]})),_:1},8,["disabled"])]})),default:h((function(){return[D,g("div",S,[b(l,{size:"small",type:"text",onClick:t[5]||(t[5]=function(e){return ae.value=!1})},{default:h((function(){return[E]})),_:1}),b(l,{size:"small",type:"primary",onClick:ue},{default:h((function(){return[T]})),_:1})])]})),_:1},8,["visible"])]),b(d,{data:Q.value,onSortChange:te,onSelectionChange:re},{default:h((function(){return[b(c,{type:"selection",width:"55"}),b(c,{align:"left",label:"id","min-width":"60",prop:"ID",sortable:"custom"}),b(c,{align:"left",label:"API路径","min-width":"150",prop:"path",sortable:"custom"}),b(c,{align:"left",label:"API分组","min-width":"150",prop:"apiGroup",sortable:"custom"}),b(c,{align:"left",label:"API简介","min-width":"150",prop:"description",sortable:"custom"}),b(c,{align:"left",label:"请求","min-width":"150",prop:"method",sortable:"custom"},{default:h((function(e){return[g("div",null,x(e.row.method)+" / "+x((t=e.row.method,n=k.value.filter((function(e){return e.value===t}))[0],n&&"".concat(n.label))),1)];var t,n})),_:1}),b(c,{align:"left",fixed:"right",label:"操作",width:"200"},{default:h((function(e){return[b(l,{icon:"edit",size:"small",type:"text",onClick:function(t){return pe(e.row)}},{default:h((function(){return[F]})),_:2},1032,["onClick"]),b(l,{icon:"delete",size:"small",type:"text",onClick:function(t){return fe(e.row)}},{default:h((function(){return[B]})),_:2},1032,["onClick"])]})),_:1})]})),_:1},8,["data"]),g("div",H,[b(V,{"current-page":W.value,"page-size":N.value,"page-sizes":[10,30,50,100],total:J.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:ee,onSizeChange:$},null,8,["current-page","page-size","total"])])]),b(C,{modelValue:ie.value,"onUpdate:modelValue":t[12]||(t[12]=function(e){return ie.value=e}),"before-close":se,title:oe.value},{footer:h((function(){return[g("div",L,[b(l,{size:"small",onClick:se},{default:h((function(){return[K]})),_:1}),b(l,{size:"small",type:"primary",onClick:de},{default:h((function(){return[M]})),_:1})])]})),default:h((function(){return[b(s,{title:"新增API，需要在角色管理内配置权限才可使用"}),b(o,{ref_key:"apiForm",ref:le,model:_.value,rules:P.value,"label-width":"80px"},{default:h((function(){return[b(r,{label:"路径",prop:"path"},{default:h((function(){return[b(n,{modelValue:_.value.path,"onUpdate:modelValue":t[8]||(t[8]=function(e){return _.value.path=e}),autocomplete:"off"},null,8,["modelValue"])]})),_:1}),b(r,{label:"请求",prop:"method"},{default:h((function(){return[b(u,{modelValue:_.value.method,"onUpdate:modelValue":t[9]||(t[9]=function(e){return _.value.method=e}),placeholder:"请选择",style:{width:"100%"}},{default:h((function(){return[(v(!0),m(y,null,w(k.value,(function(e){return v(),j(a,{key:e.value,label:"".concat(e.label,"(").concat(e.value,")"),value:e.value},null,8,["label","value"])})),128))]})),_:1},8,["modelValue"])]})),_:1}),b(r,{label:"api分组",prop:"apiGroup"},{default:h((function(){return[b(n,{modelValue:_.value.apiGroup,"onUpdate:modelValue":t[10]||(t[10]=function(e){return _.value.apiGroup=e}),autocomplete:"off"},null,8,["modelValue"])]})),_:1}),b(r,{label:"api简介",prop:"description"},{default:h((function(){return[b(n,{modelValue:_.value.description,"onUpdate:modelValue":t[11]||(t[11]=function(e){return _.value.description=e}),autocomplete:"off"},null,8,["modelValue"])]})),_:1})]})),_:1},8,["model","rules"])]})),_:1},8,["modelValue","title"])])}}});e("default",p(J,[["__scopeId","data-v-7b528490"]]))}}}))}();