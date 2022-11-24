{{if .Account}}
<div class="hidden-md hidden-lg">
    <a href="{{link `share_new_get`}}" class="btn btn-default btn-block apollo-new-share">
        创建新分享
    </a>
</div>
{{ end }}
<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			<ul class="nav nav-pills">
				<li id="tab_0">
					<a href="{{link `home_page` `t` `0`}}" style="padding: 1px 15px;">全部</a>
				</li>
				<li id="tab_9">
					<a href="{{link `tag_list`}}" style="padding: 1px 15px;">按标签</a>
				</li>
			</ul>
		</div>
		<div class="panel-body paginate-bot">
			{{range .Shares}}
				{{template "shares/_cell.tpl" .}}
			{{end}}
			<ul id="page"></ul>
		</div>
	</div>
</div>
<div class="col-md-3">
  {{template "home/_sidebar.tpl" . }}
</div>

</div>

<div class="placeholder-body"></div>

<script type="text/javascript">
  $(function () {
	$("#tab_{{.TabIndex}}").addClass("active");
    $("#page").bootstrapPaginator({
      currentPage: '{{.CurrentPage}}',
      totalPages: '{{.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        window.location.href = "/?page=" + page + "&t={{.TabIndex}}"
      }
    });
  });
</script>