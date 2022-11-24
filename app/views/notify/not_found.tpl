<div id="not-found-common" style="text-align:center">
</div>
<div class="placeholder-body"></div>

<script type="text/javascript">
    {{ if .Timeout}}
    var t = {{.Timeout}};
    {{ else }}
    var t = 15;
    {{ end }}
    function showTime() {
        $("#not-found-common").html("页面不存在😭😭😭，" + t + " 秒后自动跳转到 <a href='/'>首页</a>");
        t -= 1;
        if (t == 0) {
            window.location.href = '/';
            return;
        }
        setTimeout("showTime()",1000);
    }
    showTime();
</script>