<div class="media">
	<div class="media-left">
		<div class="cell-a-avatar">
			<a href="{{link `user_detail` `uid` .User.ID}}" style="color:white;">
				<div class="avatar">{{.User.Nickname | firstChar}}</div>
			</a>
		</div>
	</div>
	<div class="media-body">
		<div class="title">
			<a href="{{link `share_detail` `id` .ID}}">《{{.Title}}》</a>· 
			{{if .URL}}
				<span style="font-size: 14px">
					(<a href="{{link `share_direct_jump` `id` .ID}}" target="_blank">原文</a>)
				</span>
			{{else}}
				<span style="font-size: 14px">
					(<a href="{{link `share_detail` `id` .ID}}" style="color: #3c763d;">原创</a>)
				</span>
			{{end}}
		</div>
		<div class="reviews">
			评述：{{str_limit_length .Review 100}}
		</div>
		<p class="gray">
			{{year_month_str .CreatedAt}} · 推荐自 <a href="{{link `user_detail` `uid` .User.ID}}">{{.User.Nickname}}</a> · 
			<a href="{{link `tag_detail` `t` .Tag}}" class="index-share-tag">{{.Tag}}</a>
		</p>
	</div>
</div>
<div class="divide mar-top-5"></div>