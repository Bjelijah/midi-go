(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-61bb927b"],{"0d70":function(e,t,a){e.exports={mainColor:"#4090cf",navHeight:"0.8rem",navBg:"#4090cf",navText:"#4090cf",navActiveText:"#4090cf",sidebarBg:"#fff",sidebarActiveBg:"#cce5f0",sidebarText:"#67c23a",sidebarActiveText:"#67c23a",sideBarWidth:"2.8rem",sideBarWidthFold:"0.8rem",mainFontColor:"#191f25",footerHeight:"1.6rem",footerBg:"#727171"}},"1c5f":function(e,t,a){e.exports={mainColor:"#269bff",navHeight:"0.613333rem",navBg:"#269bff",navText:"#a0bdf4",navActiveText:"#fff",sidebarBg:"#fff",sidebarActiveBg:"#cce5f0",sidebarText:"#67c23a",sidebarActiveText:"#67c23a",sideBarWidth:"2.8rem",sideBarWidthFold:"0.8rem",mainFontColor:"#191f25"}},6886:function(e,t,a){"use strict";a("0d70")},"969c":function(e,t,a){"use strict";a.r(t);var s=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"page"},[a("div",{staticClass:"layout"},[a("el-form",{ref:"formRef",staticClass:"content",attrs:{"label-position":"left",model:e.formData}},[a("el-form-item",{staticStyle:{width:"100%"},attrs:{label:"分享img_url",rules:e.formRules.share_img_url,prop:"share_img_url"}},[a("el-input",{attrs:{clearable:!0,placeholder:"请填写分享图片"},model:{value:e.formData.share_img_url,callback:function(t){e.$set(e.formData,"share_img_url",t)},expression:"formData.share_img_url"}})],1),a("el-form-item",{staticStyle:{width:"100%"},attrs:{label:"title",rules:e.formRules.title,prop:"title"}},[a("el-input",{attrs:{clearable:!0,placeholder:"请填写title"},model:{value:e.formData.title,callback:function(t){e.$set(e.formData,"title",t)},expression:"formData.title"}})],1),a("el-form-item",{staticStyle:{width:"100%"},attrs:{label:"content",rules:e.formRules.content,prop:"content"}},[a("el-input",{attrs:{clearable:!0,placeholder:"请填写content"},model:{value:e.formData.content,callback:function(t){e.$set(e.formData,"content",t)},expression:"formData.content"}})],1),a("el-form-item",{staticStyle:{width:"100%"},attrs:{label:"图片压缩包",rules:e.formRules.img_zip,prop:"img_zip"}},[a("file-upload",{staticClass:"content-content",on:{success:e.zipSuccess,error:e.zipError},model:{value:e.formData.img_zip,callback:function(t){e.$set(e.formData,"img_zip",t)},expression:"formData.img_zip"}})],1)],1),a("div",{staticStyle:{height:"1px","background-color":"#DCDFE6"}}),a("div",{staticStyle:{padding:"0px 20px",margin:"20px 0px"}},[a("el-button",{staticStyle:{"margin-bottom":"10px"},attrs:{size:"small",type:"primary"},on:{click:e.doFinish}},[e._v("完成")]),a("el-input",{attrs:{disabled:!0,autosize:{minRows:3},type:"textarea"},model:{value:e.deployRes,callback:function(t){e.deployRes=t},expression:"deployRes"}})],1)],1)])},i=[],r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-upload",e._b({ref:"uploadRef",staticClass:"upload-wrapper",attrs:{headers:e.headers,data:{typeCode:e.typeCode},action:e.actionApi,"before-upload":e.beforeUpload,"on-success":e.handleSuccess,"on-error":e.handleError,"on-change":e.handleChangeFile,"on-remove":e.handleRemoveFile}},"el-upload",e.$attrs,!1),[e._t("inner",[a("el-button",{staticStyle:{"margin-bottom":"10px"},attrs:{size:"small",type:"primary",disabled:!e.uploadEnable}},[e._v("点击上传")])]),"text"!==e.listType&&e.uploadSuccess?a("div",{staticClass:"control-wrapper"},[a("label",{staticClass:"success-label"},[a("i",{staticClass:"el-icon-check"})])]):e._e(),e._t("control",["text"!==e.listType?a("div",{staticClass:"control-wrapper",on:{mouseover:e.showImageControl,mouseout:e.hideImageControl}},[a("div",{staticClass:"mask",style:{visibility:e.showControl?"visible":"hidden"}},[a("i",{staticClass:"el-icon-delete",on:{click:function(t){return t.stopPropagation(),e.handleRemoveFile(t)}}})])]):e._e()]),a("div",{staticClass:"el-upload__tip",attrs:{slot:"tip"},slot:"tip"},[e._v(e._s(e.tip))])],2)},o=[],l={name:"FileUpload",props:{typeCode:{type:String,default:"img"},listType:{type:String,default:"picture"},tip:{type:String,default:""},uploadEnable:{type:Boolean,default:!0}},data(){return{actionApi:"/api/tools/file/upload",headers:{},showControl:!1,uploadSuccess:!1}},computed:{},watch:{},mounted(){},methods:{handleSuccess(e,t){200===e.code?(this.uploadSuccess=!0,this.$emit("success",e.data.resId,e.data.resUrl)):(this.uploadSuccess=!1,this.$emit("error"),this.$message.error("文件上传失败: "+e.msg))},handleError(e){this.uploadSuccess=!1,this.$emit("error"),this.$message.error("文件上传失败: "+e)},handleRemoveFile(){this.showControl=!1,this.uploadSuccess=!1,this.$emit("remove")},handleChangeFile(e){"ready"===e.status&&(this.uploadSuccess=!1)},beforeUpload(e){if(console.log(this.$attrs),this.$attrs["size-limit"]){const t=e.size/1024/1024<this.$attrs["size-limit"];if(!t){let e=this.$attrs["size-limit"]+"MB";return this.$attrs["size-limit"]<1&&(e=1024*this.$attrs["size-limit"]+"KB"),this.$message.error("上传文件不能超过"+e),!1}}return!0},showImageControl(){const e=this.$refs.uploadRef.$el.getElementsByTagName("img")[0];e&&e.src&&(this.showControl=!0)},hideImageControl(){this.showControl=!1}}},n=l,c=(a("d6f6"),a("2877")),m=Object(c["a"])(n,r,o,!1,null,null,null),d=m.exports,p=a("bc3a"),u=a.n(p),h=a("5c96"),f=a("4360");a("bb31");const g=u.a.create({baseURL:"/api",timeout:0});g.interceptors.request.use(e=>(f["a"].getters.token&&(e.headers["Authorization"]="Bearer "+f["a"].getters.token),e),e=>(console.error(e),Promise.reject(e))),g.interceptors.response.use(e=>{const t=e.data;return 200!==t.code?(Object(h["Message"])({message:t.message,type:"error",duration:5e3}),Promise.reject(t.msg)):e.data},e=>(console.log("err"+e),Object(h["Message"])({message:e.message,type:"error",duration:5e3}),Promise.reject(e)));var b=g;function _(e){return b({url:"/tools/file/download",method:"post",data:e})}function v(e){return b({url:"/tools/autoDeploy",method:"post",data:e})}var y={name:"Index",components:{FileUpload:d},data(){return{selType:"",loading:!1,finishReady:!1,formData:{share_img_url:"",title:"",content:"",img_zip:"",img_zip_name:""},formRules:{share_img_url:[{required:!0,message:"请填写分享图片",trigger:"change"}],title:[{required:!0,message:"请填写title",trigger:"change"}],content:[{required:!0,message:"请填写title",trigger:"change"}],img_zip:[{required:!0,message:"请上传图片包",trigger:"change"}]},deployRes:""}},computed:{},methods:{download(){this.loading=!0,_({type:this.selType}).then(e=>{const t=e.data.resUrl,a=document.createElement("a");a.href=t,a.click()}).catch(e=>{})},zipSuccess(e,t){this.formData.img_zip=e,this.formData.img_zip_name=t},zipError(){this.formData.img_zip=""},doFinish(){this.$refs.formRef.validate(e=>{e&&v({img_url:this.formData.share_img_url,title:this.formData.title,content:this.formData.content,zipRes:this.formData.img_zip,name:this.formData.img_zip_name}).then(e=>{console.log("res="+JSON.stringify(e)),this.deployRes=e.data.res})})}}},C=y,x=(a("6886"),Object(c["a"])(C,s,i,!1,null,null,null));t["default"]=x.exports},d6f6:function(e,t,a){"use strict";a("1c5f")}}]);