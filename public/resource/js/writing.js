var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var cos = null;
var MdEditor = null;
var headInput = null;
var ArticleItem = {};
var cid = 1;

function setAjaxToken(xhr) {
  xhr.setRequestHeader("Authorization", localStorage.getItem("AUTH_TOKEN"));
}

function initEditor() {
  // 取默认标题
  headInput.val(ArticleItem.title);
  // 初始化编辑器
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    editorTheme: "default",
    path: CNDURL + "/lib/",
    placeholder: "",
    appendMarkdown: ArticleItem.markdown,
    codeFold: true,
    saveHTMLToTextarea: true,
    tocm: true,
    imageUpload: true,
    taskList: true,
    emoji: true,
    imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    // imageUploadURL: "/api/v1/uploadfile",

  });
}
// function uploadImage(file, cb) {
//   const config = {
//     useCdnDomain: true,
//     //自行去改七牛云的空间区域的配置
//     region: qiniu.region.z1
//   };
//   const putExtra = {
//   };
//   // 异步获取临时密钥
//   $.ajax({
//     url: "/api/v1/qiniu/token",
//     type: "GET",
//     contentType: "application/json",
//     success: function (res) {
//       if (res.code !== 200) return alert(res.error);
//       const token = res.data;
//       const observable = qiniu.upload(file, "goblog/upload/"+Date.now() + "_" + file.name, token, putExtra, config)
//       const observer = {
//         next(res){
//           // ...
//         },
//         error(err){
//           // ...
//         },
//         complete(res){
//           console.log(res)
//           cb("https://static.mszlu.com/" + res.key)
//         }
//       }
//       const subscription = observable.subscribe(observer) // 上传开始
//
//     },
//     beforeSend: setAjaxToken,
//   });
//
// }


function getArticleItem(id) {
  $.ajax({
    url: "/api/v1/article/" + id,
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code != 200) {
        initEditor();
        return alert(res.error);
      }
      ArticleItem = res.data || {};
      console.log(ArticleItem)
      initActive();
      initEditor();
    },
    beforeSend: setAjaxToken,
  });
}
function initActive() {
  $(".category li[value=" + ArticleItem.cid + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");

}
function initCache() {
  headInput = $(".header-input");
  var query = new URLSearchParams(location.search);
  var _id = query.get("id");
  // console.log(_id)
  if (_id) return getArticleItem(_id);
  // 取本地缓存
  // ArticleItem.title = window.localStorage.getItem(TITLE_KEY);
  // ArticleItem.markdown = window.localStorage.getItem(CONTENT_KEY) || "";
  // initEditor
  initEditor();
  // MdEditor.setMarkdown(ArticleItem.markdown)
  // headInput.setValue(ArticleItem.title)
  // 自动保存
  setInterval(() => saveHandler, AUTO_SAVE_TIME);
}

function saveHandler() {
  window.localStorage.setItem(TITLE_KEY, headInput.val());
  window.localStorage.setItem(CONTENT_KEY, MdEditor.getMarkdown());
}
function clearHandler() {
  window.localStorage.removeItem(TITLE_KEY);
  window.localStorage.removeItem(CONTENT_KEY);
}

// 发布
function publishHandler() {
  var cate = $(".slug-input").val()
  if (!ArticleItem.cid &&!cate) return $(".publish-tip").text("请选择分类");
  ArticleItem.title = headInput.val();
  if (!ArticleItem.title) return $(".publish-tip").text("请输入标题");
  ArticleItem.markdown = MdEditor.getMarkdown();
  if (!ArticleItem.markdown) return $(".publish-tip").text("正文");
  ArticleItem.content = MdEditor.getPreviewedHTML();

  if (!cate) {
    $.ajax({
      // Authorization:localStorage.getItem("AUTH_TOKEN"),
      url: "/api/v1/article/add",
      type: ArticleItem.id ? "PUT" : "POST",
      contentType: "application/json",
      data: JSON.stringify(ArticleItem),
      success: function (res) {
        if (res.code !== 200) return alert(res.error);
        if (ArticleItem.id) return $(".publish-tip").text("更新成功");
        ArticleItem = res.data || {};
        if (!ArticleItem.id) {
          clearHandler();
        }
        location.search = "?id=" + ArticleItem.id;
      },
      beforeSend: setAjaxToken,
    });
  }else{
    $.ajax({
      // Authorization:localStorage.getItem("AUTH_TOKEN"),
      url: "/api/v1/category/add",
      type: "POST",
      contentType: "application/json",
      data: JSON.stringify({name:cate}),
      success: function (res) {
        if (res.code !== 200) {return
        alert(res.error);}
        var info = res || {};
        var obj = JSON.parse(JSON.stringify(info.data))
        console.log(obj.id)
        cid = parseInt(obj.id)
        // console.log(cid)
        ArticleItem.cid = parseInt(cid)
        // console.log(ArticleItem.cid)
        $.ajax(
            {
              // Authorization:localStorage.getItem("AUTH_TOKEN"),
              url: "/api/v1/article/add",
              type: ArticleItem.id ? "PUT" : "POST",
              contentType: "application/json",
              data: JSON.stringify(ArticleItem),
              success: function (res) {
                if (res.code !== 200) return alert(res.error);
                if (ArticleItem.id) return $(".publish-tip").text("更新成功");
                ArticleItem = res.data || {};
                if (!ArticleItem.id) {
                  clearHandler();
                }
                location.search = "?id=" + ArticleItem.id;
              },
              beforeSend: setAjaxToken,
            });
      },
      beforeSend: setAjaxToken,
    });
    // console.log(ArticleItem.cid)



  }


}

function updatearticle(){
  //拿到文章id
  var query = new URLSearchParams(location.search);
  var id = query.get("id");
  getArticleItem(id)
  ArticleItem.markdown = MdEditor.getMarkdown()
  ArticleItem.title = headInput.val()
  ArticleItem.content=MdEditor.getPreviewedHTML()

  $.ajax(
      {// Authorization:localStorage.getItem("AUTH_TOKEN"),
        url: "/api/v1/article/"+id,
        type: "PUT",
        contentType: "application/json",
        data: JSON.stringify({ cid: ArticleItem.cid, content: ArticleItem.content , markdown:ArticleItem.markdown,title:ArticleItem.title }),
        success: function (res) {
          if (res.code !== 200) return alert(res.error);
          if (ArticleItem.id) return $(".publish-tip").text("更新成功");
          console.log(ArticleItem)
          ArticleItem = res.data || {};
          if (!ArticleItem.id) {
            clearHandler();
          }
        },
        beforeSend: setAjaxToken,
      });
  location.reload()
}

$(function () {
  // 初始化缓存
  initCache();
  // 返回首页
  console.log("1")
  var back = $(".home-btn");
  back.click(function () {
    clearHandler()
    location.href = ArticleItem.id ? "/article/" + ArticleItem.id + ".html" : "/";
  });
  if (location.search) back.text("查看");
  // 保存
  $(".save-btn").click(saveHandler);
  var drop = $(".publish-drop");
  // 显示
  $(".publish-show").click(function () {
    drop.show();
  });
  // 隐藏
  $(".publish-close").click(function () {
    drop.hide();
  });
  $(".cancel-btn").click(function () {
    drop.hide();
  });
  // 发布逻辑
  if (!location.search) {
    $(".publish-btn").click(publishHandler);
  } else{
    $(".publish-btn").click(updatearticle)
  }
  // 选择分类
  $(".category").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.cid = parseInt(target.attr("value"));
    $(".publish-tip").text("");
  });

});
