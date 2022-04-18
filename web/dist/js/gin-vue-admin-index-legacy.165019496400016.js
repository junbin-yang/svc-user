/*! 
 Build Time : 1650194964000 */
!function(){function e(e,t,a,n,l,r,u){try{var i=e[r](u),o=i.value}catch(c){return void a(c)}i.done?t(o):Promise.resolve(o).then(n,l)}function t(t){return function(){var a=this,n=arguments;return new Promise((function(l,r){var u=t.apply(a,n);function i(t){e(u,l,r,i,o,"next",t)}function o(t){e(u,l,r,i,o,"throw",t)}i(void 0)}))}}var a=document.createElement("style");a.innerHTML=".previewCodeTool[data-v-143ab70e]{display:flex;align-items:center;padding:5px 0}.button-box[data-v-143ab70e]{padding:10px 20px}.button-box .el-button[data-v-143ab70e]{margin-right:20px;float:right}.auto-btn-list[data-v-143ab70e]{margin-top:16px}.auto-icon[data-v-143ab70e]{margin-left:6px;color:#666;cursor:pointer}\n",document.head.appendChild(a),System.register(["./gin-vue-admin-fieldDialog-legacy.1650194964000.js","./gin-vue-admin-previewCodeDialg-legacy.1650194964000.js","./gin-vue-admin-stringFun-legacy.1650194964000.js","./gin-vue-admin-autoCode-legacy.1650194964000.js","./gin-vue-admin-dictionary-legacy.1650194964000.js","../gva/gin-vue-admin-index-legacy.1650194964000.js","./gin-vue-admin-sysDictionary-legacy.1650194964000.js","./gin-vue-admin-warningBar-legacy.1650194964000.js"],(function(e){"use strict";var a,n,l,r,u,i,o,c,d,f,s,p,m,v,b,g,y,x,h,N,k,w,_,C,V,T,R,U,S,z,j,F,D,L,q,J;return{setters:[function(e){a=e.default},function(e){n=e.default},function(e){l=e.a,r=e.b,u=e.t,i=e.c},function(e){o=e.p,c=e.c,d=e.g,f=e.a,s=e.b,p=e.d,m=e.e},function(e){v=e.u},function(e){b=e._,g=e.J,y=e.u,x=e.a,h=e.r,N=e.b,k=e.o,w=e.c,_=e.d,C=e.e,V=e.w,T=e.G,R=e.H,U=e.g,S=e.ah,z=e.l,j=e.f,F=e.i,D=e.p,L=e.q,q=e.j,J=e.ai},function(){},function(){}],execute:function(){var A=function(){var e=t(regeneratorRuntime.mark((function e(t){var a;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return a=v(),e.next=3,a.getDictionary(t);case 3:return e.abrupt("return",a.dictionaryMap[t]);case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),O=function(e){return D("data-v-143ab70e"),e=e(),L(),e},E={class:"gva-search-box"},P={style:{fontSize:"16px",paddingLeft:"20px"}},M=F(" 点这里从现有数据库创建代码 "),$=F("使用此表创建"),B={class:"gva-search-box"},I=O((function(){return _("div",null," 自动创建API ",-1)})),X=O((function(){return _("div",null," 自动移动文件 ",-1)})),H={class:"gva-table-box"},G={class:"gva-btn-list"},K=F("新增Field"),Q=F("编辑"),W=F("上移"),Y=F("下移"),Z=O((function(){return _("p",null,"确定删除吗？",-1)})),ee={style:{"text-align":"right","margin-top":"8px"}},te=F("取消"),ae=F("确定"),ne=F("删除"),le={class:"gva-btn-list justify-content-flex-end auto-btn-list"},re=F("预览代码"),ue=F("生成代码"),ie={class:"dialog-footer"},oe=F("取 消"),ce=F("确 定"),de={class:"previewCodeTool"},fe=O((function(){return _("p",null,"操作栏：",-1)})),se=F("全选"),pe=F("复制"),me={class:"dialog-footer",style:{"padding-top":"14px","padding-right":"14px"}},ve=F("确 定"),be={name:"AutoCode"},ge=Object.assign(be,{setup:function(e){var v={fieldName:"",fieldDesc:"",fieldType:"",dataType:"",fieldJson:"",columnName:"",dataTypeLong:"",comment:"",fieldSearchType:"",dictType:""},b=g(),F=y(),D=x([]),L=h({}),O=h({dbName:"",tableName:""}),be=h([]),ge=h([]),ye=h(""),xe=h({}),he=h({structName:"",tableName:"",packageName:"",package:"",abbreviation:"",description:"",autoCreateApiToSql:!1,autoMoveFile:!1,fields:[]}),Ne=h({structName:[{required:!0,message:"请输入结构体名称",trigger:"blur"}],abbreviation:[{required:!0,message:"请输入结构体简称",trigger:"blur"}],description:[{required:!0,message:"请输入结构体描述",trigger:"blur"}],packageName:[{required:!0,message:"文件名称：sysXxxxXxxx",trigger:"blur"}],package:[{required:!0,message:"请选择package",trigger:"blur"}]}),ke=h({}),we=h({}),_e=h(!1),Ce=h(!1),Ve=h(null),Te=function(){Ve.value.selectText()},Re=function(){Ve.value.copy()},Ue=function(e){_e.value=!0,e?(ye.value="edit",we.value=JSON.parse(JSON.stringify(e)),ke.value=e):(ye.value="add",ke.value=JSON.parse(JSON.stringify(v)))},Se=J(),ze=function(){Se.refs.fieldDialogNode.fieldDialogFrom.validate((function(e){if(!e)return!1;ke.value.fieldName=r(ke.value.fieldName),"add"===ye.value&&he.value.fields.push(ke.value),_e.value=!1}))},je=function(){"edit"===ye.value&&(ke.value=we.value),_e.value=!1},Fe=h(null),De=function(){var e=t(regeneratorRuntime.mark((function e(a){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!(he.value.fields.length<=0)){e.next=3;break}return q({type:"error",message:"请填写至少一个field"}),e.abrupt("return",!1);case 3:if(!he.value.fields.some((function(e){return e.fieldName===he.value.structName}))){e.next=6;break}return q({type:"error",message:"存在与结构体同名的字段"}),e.abrupt("return",!1);case 6:Fe.value.validate(function(){var e=t(regeneratorRuntime.mark((function e(t){var n,i,d,f,s,p,m,v;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!t){e.next=33;break}for(n in he.value)"string"==typeof he.value[n]&&(he.value[n]=he.value[n].trim());if(he.value.structName=r(he.value.structName),he.value.tableName=he.value.tableName.replace(" ",""),he.value.tableName||(he.value.tableName=u(l(he.value.structName))),he.value.structName!==he.value.abbreviation){e.next=8;break}return q({type:"error",message:"structName和struct简称不能相同"}),e.abrupt("return",!1);case 8:if(he.value.humpPackageName=u(he.value.packageName),!a){e.next=17;break}return e.next=12,o(he.value);case 12:i=e.sent,L.value=i.data.autoCode,Ce.value=!0,e.next=31;break;case 17:return e.next=19,c(he.value);case 19:if(f=e.sent,"false"!==(null===(d=f.headers)||void 0===d?void 0:d.success)){e.next=24;break}return e.abrupt("return");case 24:if(!he.value.autoMoveFile){e.next=27;break}return q({type:"success",message:"自动化代码创建成功，自动移动成功"}),e.abrupt("return");case 27:q({type:"success",message:"自动化代码创建成功，正在下载"});case 28:s=new Blob([f]),p="ginvueadmin.zip","download"in document.createElement("a")?(m=window.URL.createObjectURL(s),(v=document.createElement("a")).style.display="none",v.href=m,v.setAttribute("download",p),document.body.appendChild(v),v.click(),document.body.removeChild(v),window.URL.revokeObjectURL(m)):window.navigator.msSaveBlob(s,p);case 31:e.next=34;break;case 33:return e.abrupt("return",!1);case 34:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}());case 7:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),Le=function(){var e=t(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,p();case 2:0===(t=e.sent).code&&(be.value=t.data.dbs);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),qe=function(){var e=t(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,d({dbName:O.value.dbName});case 2:0===(t=e.sent).code&&(ge.value=t.data.tables),O.value.tableName="";case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Je=function(){var e=t(regeneratorRuntime.mark((function e(){var t,a,n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=["id","created_at","updated_at","deleted_at"],e.next=3,f(O.value);case 3:0===(a=e.sent).code&&(n=i(O.value.tableName),he.value.structName=r(n),he.value.tableName=O.value.tableName,he.value.packageName=n,he.value.abbreviation=n,he.value.description=n+"表",he.value.autoCreateApiToSql=!0,he.value.fields=[],a.data.columns&&a.data.columns.forEach((function(e){if(!t.some((function(t){return t===e.columnName}))){var a=i(e.columnName);he.value.fields.push({fieldName:r(a),fieldDesc:e.columnComment||a+"字段",fieldType:xe.value[e.dataType],dataType:e.dataType,fieldJson:a,dataTypeLong:e.dataTypeLong&&e.dataTypeLong.split(",")[0],columnName:e.columnName,comment:e.columnComment,fieldSearchType:"",dictType:""})}})));case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Ae=function(){var e=t(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:["string","int","bool","float64","time.Time"].forEach(function(){var e=t(regeneratorRuntime.mark((function e(t){var a;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,A(t);case 2:(a=e.sent)&&a.forEach((function(e){xe.value[e.label]=t}));case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}());case 2:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Oe=function(){var e=t(regeneratorRuntime.mark((function e(t){var a;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,m({id:Number(t)});case 2:0===(a=e.sent).code&&(he.value=JSON.parse(a.data.meta));case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),Ee=h([]),Pe=function(){var e=t(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,s();case 2:0===(t=e.sent).code&&(Ee.value=t.data.pkgs);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Me=function(){F.push({name:"autoPkg"})};return function(){Le(),Ae(),Pe();var e=b.params.id;e&&Oe(e)}(),function(e,t){var r=N("pointer"),u=N("el-icon"),i=N("el-option"),o=N("el-select"),c=N("el-form-item"),d=N("el-button"),f=N("el-form"),s=N("el-collapse-item"),p=N("el-collapse"),m=N("el-input"),v=N("refresh"),b=N("document-add"),g=N("el-tooltip"),y=N("el-checkbox"),x=N("el-table-column"),h=N("el-popover"),F=N("el-table"),q=N("el-dialog");return k(),w("div",null,[_("div",E,[C(p,{modelValue:U(D),"onUpdate:modelValue":t[2]||(t[2]=function(e){return S(D)?D.value=e:function(e){throw new TypeError('"'+e+'" is read-only')}("activeNames")}),style:{"margin-bottom":"12px"}},{default:V((function(){return[C(s,{name:"1"},{title:V((function(){return[_("div",P,[M,C(u,{class:"header-icon"},{default:V((function(){return[C(r)]})),_:1})])]})),default:V((function(){return[C(f,{ref:"getTableForm",style:{"margin-top":"24px"},inline:!0,model:O.value,"label-width":"120px"},{default:V((function(){return[C(c,{label:"数据库名",prop:"structName"},{default:V((function(){return[C(o,{modelValue:O.value.dbName,"onUpdate:modelValue":t[0]||(t[0]=function(e){return O.value.dbName=e}),filterable:"",placeholder:"请选择数据库",onChange:qe},{default:V((function(){return[(k(!0),w(T,null,R(be.value,(function(e){return k(),z(i,{key:e.database,label:e.database,value:e.database},null,8,["label","value"])})),128))]})),_:1},8,["modelValue"])]})),_:1}),C(c,{label:"表名",prop:"structName"},{default:V((function(){return[C(o,{modelValue:O.value.tableName,"onUpdate:modelValue":t[1]||(t[1]=function(e){return O.value.tableName=e}),disabled:!O.value.dbName,filterable:"",placeholder:"请选择表"},{default:V((function(){return[(k(!0),w(T,null,R(ge.value,(function(e){return k(),z(i,{key:e.tableName,label:e.tableName,value:e.tableName},null,8,["label","value"])})),128))]})),_:1},8,["modelValue","disabled"])]})),_:1}),C(c,null,{default:V((function(){return[C(d,{size:"small",type:"primary",onClick:Je},{default:V((function(){return[$]})),_:1})]})),_:1})]})),_:1},8,["model"])]})),_:1})]})),_:1},8,["modelValue"])]),_("div",B,[C(f,{ref_key:"autoCodeForm",ref:Fe,rules:Ne.value,model:he.value,"label-width":"120px",inline:!0},{default:V((function(){return[C(c,{label:"Struct名称",prop:"structName"},{default:V((function(){return[C(m,{modelValue:he.value.structName,"onUpdate:modelValue":t[3]||(t[3]=function(e){return he.value.structName=e}),placeholder:"首字母自动转换大写"},null,8,["modelValue"])]})),_:1}),C(c,{label:"TableName",prop:"tableName"},{default:V((function(){return[C(m,{modelValue:he.value.tableName,"onUpdate:modelValue":t[4]||(t[4]=function(e){return he.value.tableName=e}),placeholder:"指定表名（非必填）"},null,8,["modelValue"])]})),_:1}),C(c,{label:"Struct简称",prop:"abbreviation"},{default:V((function(){return[C(m,{modelValue:he.value.abbreviation,"onUpdate:modelValue":t[5]||(t[5]=function(e){return he.value.abbreviation=e}),placeholder:"简称会作为入参对象名和路由group"},null,8,["modelValue"])]})),_:1}),C(c,{label:"Struct中文名称",prop:"description"},{default:V((function(){return[C(m,{modelValue:he.value.description,"onUpdate:modelValue":t[6]||(t[6]=function(e){return he.value.description=e}),placeholder:"中文描述作为自动api描述"},null,8,["modelValue"])]})),_:1}),C(c,{label:"文件名称",prop:"packageName"},{default:V((function(){return[C(m,{modelValue:he.value.packageName,"onUpdate:modelValue":t[7]||(t[7]=function(e){return he.value.packageName=e}),placeholder:"生成文件的默认名称(建议为驼峰格式,首字母小写,如sysXxxXxxx)",onBlur:t[8]||(t[8]=function(e){return function(e,t){e[t]=l(e[t])}(he.value,"packageName")})},null,8,["modelValue"])]})),_:1}),C(c,{label:"Package（包）",prop:"packageName"},{default:V((function(){return[C(o,{modelValue:he.value.package,"onUpdate:modelValue":t[9]||(t[9]=function(e){return he.value.package=e}),style:{width:"194px"}},{default:V((function(){return[(k(!0),w(T,null,R(Ee.value,(function(e){return k(),z(i,{key:e.ID,value:e.packageName,label:e.packageName},null,8,["value","label"])})),128))]})),_:1},8,["modelValue"]),C(u,{class:"auto-icon",onClick:Pe},{default:V((function(){return[C(v)]})),_:1}),C(u,{class:"auto-icon",onClick:Me},{default:V((function(){return[C(b)]})),_:1})]})),_:1}),C(c,null,{label:V((function(){return[C(g,{content:"注：把自动生成的API注册进数据库",placement:"bottom",effect:"light"},{default:V((function(){return[I]})),_:1})]})),default:V((function(){return[C(y,{modelValue:he.value.autoCreateApiToSql,"onUpdate:modelValue":t[10]||(t[10]=function(e){return he.value.autoCreateApiToSql=e})},null,8,["modelValue"])]})),_:1}),C(c,null,{label:V((function(){return[C(g,{content:"注：自动迁移生成的文件到ymal配置的对应位置",placement:"bottom",effect:"light"},{default:V((function(){return[X]})),_:1})]})),default:V((function(){return[C(y,{modelValue:he.value.autoMoveFile,"onUpdate:modelValue":t[11]||(t[11]=function(e){return he.value.autoMoveFile=e})},null,8,["modelValue"])]})),_:1})]})),_:1},8,["rules","model"])]),_("div",H,[_("div",G,[C(d,{size:"small",type:"primary",onClick:t[12]||(t[12]=function(e){return Ue()})},{default:V((function(){return[K]})),_:1})]),C(F,{data:he.value.fields},{default:V((function(){return[C(x,{align:"left",type:"index",label:"序列",width:"100"}),C(x,{align:"left",prop:"fieldName",label:"Field名"}),C(x,{align:"left",prop:"fieldDesc",label:"中文名"}),C(x,{align:"left",prop:"fieldJson",label:"FieldJson"}),C(x,{align:"left",prop:"fieldType",label:"Field数据类型",width:"130"}),C(x,{align:"left",prop:"dataTypeLong",label:"数据库字段长度",width:"130"}),C(x,{align:"left",prop:"columnName",label:"数据库字段",width:"130"}),C(x,{align:"left",prop:"comment",label:"数据库字段描述",width:"130"}),C(x,{align:"left",prop:"fieldSearchType",label:"搜索条件",width:"130"}),C(x,{align:"left",prop:"dictType",label:"字典",width:"130"}),C(x,{align:"left",label:"操作",width:"300"},{default:V((function(e){return[C(d,{size:"small",type:"text",icon:"edit",onClick:function(t){return Ue(e.row)}},{default:V((function(){return[Q]})),_:2},1032,["onClick"]),C(d,{size:"small",type:"text",disabled:0===e.$index,onClick:function(t){return function(e){if(0!==e){var t=he.value.fields[e-1];he.value.fields.splice(e-1,1),he.value.fields.splice(e,0,t)}}(e.$index)}},{default:V((function(){return[W]})),_:2},1032,["disabled","onClick"]),C(d,{size:"small",type:"text",disabled:e.$index+1===he.value.fields.length,onClick:function(t){return function(e){if(e!==he.value.fields.length-1){var t=he.value.fields[e+1];he.value.fields.splice(e+1,1),he.value.fields.splice(e,0,t)}}(e.$index)}},{default:V((function(){return[Y]})),_:2},1032,["disabled","onClick"]),C(h,{visible:e.row.visible,"onUpdate:visible":function(t){return e.row.visible=t},placement:"top"},{reference:V((function(){return[C(d,{size:"small",type:"text",icon:"delete",onClick:function(t){return e.row.visible=!0}},{default:V((function(){return[ne]})),_:2},1032,["onClick"])]})),default:V((function(){return[Z,_("div",ee,[C(d,{size:"small",type:"text",onClick:function(t){return e.row.visible=!1}},{default:V((function(){return[te]})),_:2},1032,["onClick"]),C(d,{type:"primary",size:"small",onClick:function(t){return function(e){he.value.fields.splice(e,1)}(e.$index)}},{default:V((function(){return[ae]})),_:2},1032,["onClick"])])]})),_:2},1032,["visible","onUpdate:visible"])]})),_:1})]})),_:1},8,["data"]),_("div",le,[C(d,{size:"small",type:"primary",onClick:t[13]||(t[13]=function(e){return De(!0)})},{default:V((function(){return[re]})),_:1}),C(d,{size:"small",type:"primary",onClick:t[14]||(t[14]=function(e){return De(!1)})},{default:V((function(){return[ue]})),_:1})])]),C(q,{modelValue:_e.value,"onUpdate:modelValue":t[15]||(t[15]=function(e){return _e.value=e}),title:"组件内容"},{footer:V((function(){return[_("div",ie,[C(d,{size:"small",onClick:je},{default:V((function(){return[oe]})),_:1}),C(d,{size:"small",type:"primary",onClick:ze},{default:V((function(){return[ce]})),_:1})])]})),default:V((function(){return[_e.value?(k(),z(a,{key:0,ref:"fieldDialogNode","dialog-middle":ke.value},null,8,["dialog-middle"])):j("",!0)]})),_:1},8,["modelValue"]),C(q,{modelValue:Ce.value,"onUpdate:modelValue":t[17]||(t[17]=function(e){return Ce.value=e})},{title:V((function(){return[_("div",de,[fe,C(d,{size:"small",type:"primary",onClick:Te},{default:V((function(){return[se]})),_:1}),C(d,{size:"small",type:"primary",onClick:Re},{default:V((function(){return[pe]})),_:1})])]})),footer:V((function(){return[_("div",me,[C(d,{size:"small",type:"primary",onClick:t[16]||(t[16]=function(e){return Ce.value=!1})},{default:V((function(){return[ve]})),_:1})])]})),default:V((function(){return[Ce.value?(k(),z(n,{key:0,ref_key:"previewNode",ref:Ve,"preview-code":L.value},null,8,["preview-code"])):j("",!0)]})),_:1},8,["modelValue"])])}}});e("default",b(ge,[["__scopeId","data-v-143ab70e"]]))}}}))}();
