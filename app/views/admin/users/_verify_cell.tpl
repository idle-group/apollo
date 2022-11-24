<div class="media">
	<div class="media-body">
		<div class="title">
			<a href="{{link `user_detail` `uid` .ID}}">{{.StuId}}</a>
		</div>
		<div class="gray">
			<span>邮箱: {{.Email}}</span> ·
			<span>权限: {{.Priority}}</span>
			<image>学生证: {{.StuPhoto}}</image>
			| <a href="{{link `admin_valid_stu_post` `uid` .ID}}">通过审核</a>
		</div>
	</div>
</div>
<div class="divide mar-top-5"></div>