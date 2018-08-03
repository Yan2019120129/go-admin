// Code generated by hero.
// source: /Users/chenhg5/go/src/goAdmin/resources/pages/base_render.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"goAdmin/auth"
	"goAdmin/menu"

	"github.com/shiyanhui/hero"
)

func BaseContent(content string, menuList []menu.MenuItem, title string, description string, user auth.User, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>GoAdmin</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <!-- Bootstrap 3.3.7 -->
    <link rel="stylesheet" href="../../assets/bootstrap/dist/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="../../assets/font-awesome/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="../../assets/Ionicons/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="../../assets/datatables.net-bs/css/dataTables.bootstrap.min.css">
    <!-- iCheck -->
    <link rel="stylesheet" href="../../assets/iCheck/minimal/_all.css">
    <link rel="stylesheet" href="../../assets/iCheck/futurico/futurico.css">
    <link rel="stylesheet" href="../../assets/iCheck/polaris/polaris.css">
    <link rel="stylesheet" href="../../assets/toastr/build/toastr.min.css">
    <link rel="stylesheet" href="../../assets/nprogress/nprogress.css">
    <link rel="stylesheet" href="../../assets/sweetalert/dist/sweetalert.css">
    <link rel="stylesheet" href="../../assets/select2/select2.min.css">
    <link rel="stylesheet" href="../../assets/fileinput/fileinput.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="../../assets/dist/css/AdminLTE.min.css">
    <!-- AdminLTE Skins. Choose a skin from the css/skins
         folder instead of downloading all of them to reduce the load. -->
    <link rel="stylesheet" href="../../assets/dist/css/skins/_all-skins.min.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="../../assets/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="../../assets/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Google Font -->
    <link rel="stylesheet"
          href="../../assets/googleapis/font.css">
</head>
<!-- skin-black skin-green-light skin-green skin-red skin-yellow skin-purple skin-purple-light -->
<body class="hold-transition skin-black sidebar-mini">
<div class="wrapper">

    <header class="main-header">
        <!-- Logo -->
        <a href="../../index2.html" class="logo">
            <!-- mini logo for sidebar mini 50x50 pixels -->
            <span class="logo-mini"><b>G</b>A</span>
            <!-- logo for regular state and mobile devices -->
            <span class="logo-lg"><b>Go</b>Admin</span>
        </a>
        <!-- Header Navbar: style can be found in header.less -->
        <nav class="navbar navbar-static-top">
            <!-- Sidebar toggle button-->
            <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </a>

            <div class="navbar-custom-menu">
                <ul class="nav navbar-nav">
                    <!-- User Account: style can be found in dropdown.less -->
                    <li class="dropdown user user-menu">
                        `)
	buffer.WriteString(`
<a href="#" class="dropdown-toggle" data-toggle="dropdown">
    <img src="http://localhost:4003`)
	hero.EscapeHTML(user.Avatar, buffer)
	buffer.WriteString(`" class="user-image" alt="User Image">
    <span class="hidden-xs">`)
	hero.EscapeHTML(user.Name, buffer)
	buffer.WriteString(`</span>
</a>
<ul class="dropdown-menu">
    <!-- User image -->
    <li class="user-header">
        <img src="http://localhost:4003`)
	hero.EscapeHTML(user.Avatar, buffer)
	buffer.WriteString(`" class="img-circle" alt="User Image">
        <p>
            `)
	hero.EscapeHTML(user.Name, buffer)
	buffer.WriteString(` - `)
	hero.EscapeHTML(user.LevelName, buffer)
	buffer.WriteString(`
            <small>`)
	hero.EscapeHTML(user.CreateAt, buffer)
	buffer.WriteString(`</small>
        </p>
    </li>
    <li class="user-footer">
        <div class="pull-left">
            <a href="#" class="btn btn-default btn-flat">Setting</a>
        </div>
        <div class="pull-right">
            <a href="/logout" class="btn btn-default btn-flat">Sign out</a>
        </div>
    </li>
</ul>
`)

	buffer.WriteString(`
                    </li>
                </ul>
            </div>
        </nav>
    </header>
    <!-- Left side column. contains the logo and sidebar -->
    <aside class="main-sidebar">
        <!-- sidebar: style can be found in sidebar.less -->
        <section class="sidebar">
            <!-- Sidebar user panel -->
            <div class="user-panel">
                <div class="pull-left image">
                    <img src="../../assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                </div>
                <div class="pull-left info">
                    `)
	buffer.WriteString(`
<p>`)
	hero.EscapeHTML(user.Name, buffer)
	buffer.WriteString(`</p>
`)

	buffer.WriteString(`
                    <a href="#"><i class="fa fa-circle text-success"></i> Online</a>
                </div>
            </div>
            <!-- search form -->
            <form action="#" method="get" class="sidebar-form">
                <div class="input-group">
                    <input type="text" name="q" class="form-control" placeholder="Search...">
                    <span class="input-group-btn">
                <button type="submit" name="search" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
                </button>
              </span>
                </div>
            </form>
            <!-- /.search form -->
            <!-- sidebar menu: : style can be found in sidebar.less -->
            <ul class="sidebar-menu" data-widget="tree">
                <li class="header">MAIN NAVIGATION</li>
                `)
	for _, list := range menuList {
		if len(list.ChildrenList) == 0 {
			buffer.WriteString(`
<li class='`)
			hero.EscapeHTML(list.Active, buffer)
			buffer.WriteString(`'>
    <a href='`)
			hero.EscapeHTML(list.Url, buffer)
			buffer.WriteString(`'>
        <i class="fa `)
			hero.EscapeHTML(list.Icon, buffer)
			buffer.WriteString(`"></i><span>`)
			hero.EscapeHTML(list.Name, buffer)
			buffer.WriteString(`</span>
        <span class="pull-right-container"><!-- <small class="label pull-right bg-green">new</small> --></span>
    </a>
</li>
`)
		} else {
			buffer.WriteString(`
<li class="treeview `)
			hero.EscapeHTML(list.Active, buffer)
			buffer.WriteString(`">
    <a href="#">
        <i class="fa `)
			hero.EscapeHTML(list.Icon, buffer)
			buffer.WriteString(`"></i> <span>`)
			hero.EscapeHTML(list.Name, buffer)
			buffer.WriteString(`</span>
        <span class="pull-right-container">
                                  <i class="fa fa-angle-left pull-right"></i>
                                </span>
    </a>
    <ul class="treeview-menu">
        `)
			for _, item := range list.ChildrenList {
				buffer.WriteString(`
        <li><a href="`)
				hero.EscapeHTML(item.Url, buffer)
				buffer.WriteString(`"><i class="fa `)
				hero.EscapeHTML(item.Icon, buffer)
				buffer.WriteString(`"></i> `)
				hero.EscapeHTML(item.Name, buffer)
				buffer.WriteString(`</a></li>
        `)
			}
			buffer.WriteString(`
    </ul>
</li>
`)
		}
	}

	buffer.WriteString(`
            </ul>
        </section>
        <!-- /.sidebar -->
    </aside>

    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper" id="pjax-container">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            `)
	buffer.WriteString(`
<h1>
    `)
	hero.EscapeHTML(title, buffer)
	buffer.WriteString(`
    <small>`)
	hero.EscapeHTML(description, buffer)
	buffer.WriteString(`</small>
</h1>
`)

	buffer.WriteString(`
        </section>

        <!-- Main content -->
        <section class="content">
            `)
	buffer.WriteString(content)

	buffer.WriteString(`
        </section>
        <!-- /.content -->
    </div>
    <!-- /.content-wrapper -->
    <footer class="main-footer">
        <div class="pull-right hidden-xs">
            <b>Version</b> 0.0.1
        </div>
        <strong>Copyright &copy; 2018- <a href="https://github.com/chenhg5/go-admin">GoAdmin</a>.</strong> All rights
        reserved.
    </footer>
</div>
<!-- ./wrapper -->

<!-- jQuery 3 -->
<!-- <script src="../../assets/jquery/dist/jquery.min.js"></script> -->
<script src="../../assets/jQuery/jQuery-2.1.4.min.js"></script>
<!-- Bootstrap 3.3.7 -->
<script src="../../assets/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="../../assets/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="../../assets/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="../../assets/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="../../assets/fastclick/lib/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="../../assets/dist/js/adminlte.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="../../assets/dist/js/demo.js"></script>
<script src="../../assets/select2/select2.full.min.js"></script>
<script src="../../assets/fileinput/fileinput.min.js"></script>
<script src="../../assets/iCheck/icheck.min.js"></script>
<script src="../../assets/nprogress/nprogress.js"></script>
<script src="../../assets/toastr/build/toastr.min.js"></script>
<script src="../../assets/bootstrap3-editable/js/bootstrap-editable.min.js"></script>
<script src="../../assets/jquery-pjax/jquery.pjax.js"></script>
<script src="../../assets/sweetalert/dist/sweetalert.min.js"></script>
<script>
    $('.grid-per-pager').on("change", function (e) {
        console.log("changing...")
        $.pjax({url: this.value, container: '#pjax-container'});
    });
    $('.grid-refresh').on('click', function () {
        $.pjax.reload('#pjax-container');
        toastr.success('Refresh succeeded !');
    });
    // edit result notify
    // toastr.success('Refresh succeeded !');
    $.fn.editable.defaults.params = function (params) {
        params._token = LA.token;
        params._editable = 1;
        params._method = 'PUT';
        return params;
    };

    $.fn.editable.defaults.error = function (data) {
        var msg = '';
        if (data.responseJSON.errors) {
            $.each(data.responseJSON.errors, function (k, v) {
                msg += v + "\n";
            });
        }
        return msg
    };

    toastr.options = {
        closeButton: true,
        progressBar: true,
        showMethod: 'slideDown',
        timeOut: 4000
    };

    $.pjax.defaults.timeout = 5000;
    $.pjax.defaults.maxCacheLength = 0;
    $(document).pjax('a:not(a[target="_blank"])', {
        container: '#pjax-container'
    });

    NProgress.configure({parent: '#pjax-container'});

    $(document).on('pjax:timeout', function (event) {
        event.preventDefault();
    });

    $(document).on('submit', 'form[pjax-container]', function (event) {
        $.pjax.submit(event, '#pjax-container')
    });

    $(document).on("pjax:popstate", function () {

        $(document).one("pjax:end", function (event) {
            $(event.target).find("script[data-exec-on-popstate]").each(function () {
                $.globalEval(this.text || this.textContent || this.innerHTML || '');
            });
        });
    });

    $(document).on('pjax:send', function (xhr) {
        if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
            $submit_btn = $('form[pjax-container] :submit');
            if ($submit_btn) {
                $submit_btn.button('loading')
            }
        }
        NProgress.start();
    });

    $(document).on('pjax:complete', function (xhr) {
        if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
            $submit_btn = $('form[pjax-container] :submit');
            if ($submit_btn) {
                $submit_btn.button('reset')
            }
        }
        NProgress.done();
    });

    $(function () {
        $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
            var $parent = $(this).parent().addClass('active');
            $parent.siblings('.treeview.active').find('> a').trigger('click');
            $parent.siblings().removeClass('active').find('li').removeClass('active');
        });

        $('[data-toggle="popover"]').popover();
    });

    // (function ($) {
    //     $.fn.admin = LA;
    //     $.admin = LA;
    //
    // })(jQuery);

    $('.grid-row-delete').unbind('click').click(function () {

        var id = $(this).data('id');

        swal({
                    title: "你确定要删除吗",
                    type: "warning",
                    showCancelButton: true,
                    confirmButtonColor: "#DD6B55",
                    confirmButtonText: "确定",
                    closeOnConfirm: false,
                    cancelButtonText: "取消"
                },
                function () {
                    $.ajax({
                        method: 'post',
                        url: '/user/delete',
                        data: {
                            id: id
                        },
                        success: function (data) {
                            $.pjax.reload('#pjax-container');

                            if (typeof data === 'object') {
                                if (data.status) {
                                    swal(data.message, '', 'success');
                                } else {
                                    swal(data.message, '', 'error');
                                }
                            }
                        },
                        error: function (data) {
                            swal("删除失败", '', 'error');
                        }
                    });
                });
    });

    $('.grid-select-all').on('ifChanged', function (event) {
        if (this.checked) {
            $('.grid-row-checkbox').iCheck('check');
        } else {
            $('.grid-row-checkbox').iCheck('uncheck');
        }
    });
    $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'});

    $(function () {
        $('.grid-row-checkbox').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
            if (this.checked) {
                $(this).closest('tr').css('background-color', '#ffffd5');
            } else {
                $(this).closest('tr').css('background-color', '');
            }
        });
    });
</script>
</body>
</html>
`)

}
