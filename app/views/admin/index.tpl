<div class="row">
  <div class="col-md-9 mt-3"> 
    <div class="panel panel-default">
      <div class="panel-body paginate-bot">
        <ul>
          <li><a href="{{link `admin_account_list`}}">用户列表</a></li>
          <li><a href="{{link `admin_comments_list`}}">评论列表</a></li>
          <li><a href="{{link `admin_valid_stu`}}">实名审核</a></li>
        </ul>
      </div>
      </div>
  </div>
  <div class="col-md-3">
    {{template "home/_sidebar.tpl" . }}
  </div>
</div>