
<!doctype html>
<html lang="en">

<head>
    <meta charset="utf-8" />
    <link rel="apple-touch-icon" sizes="76x76" href="../assets/img/apple-icon.png" />
    <link rel="icon" type="image/png" href="../assets/img/favicon.png" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <title>锁定</title>
    <meta content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0' name='viewport' />
    <meta name="viewport" content="width=device-width" />
    <!-- Bootstrap core CSS     -->
    <link href="/static/assets/css/bootstrap.min.css" rel="stylesheet" />
    <!--  Material Dashboard CSS    -->
    <link href="/static/assets/css/material-dashboard.css" rel="stylesheet" />
    <!--  CSS for Demo Purpose, don't include it in your project     -->
    <link href="/static/assets/css/demo.css" rel="stylesheet" />
    <!--     Fonts and icons     -->
    <!--<link href="http://maxcdn.bootstrapcdn.com/font-awesome/latest/css/font-awesome.min.css" rel="stylesheet">-->
    <link href="/static/assets/css/font-awesome.min.css" rel="stylesheet">
    <link rel="stylesheet" type="text/css" href="static/assets/css/icons.css" />
    <style>
        .fixed-plugin .dropdown-menu:before, .fixed-plugin .dropdown-menu:after {
            top: 43px;
        }
    </style>
</head>

<body>
<nav class="navbar navbar-primary navbar-transparent navbar-absolute">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navigation-example-2">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="javascript:void(0)">Night Cloud</a>
        </div>
        <div class="collapse navbar-collapse">
            <ul class="nav navbar-nav navbar-right">
                <li>
                    <a href="/console">
                        <i class="material-icons">dashboard</i> 控制台
                    </a>
                </li>
                <li class=" active ">
                    <a href="javascript:void(0)">
                        <i class="material-icons">lock_open</i> Lock
                    </a>
                </li>
            </ul>
        </div>
    </div>
</nav>
<div class="wrapper wrapper-full-page">
    <div class="full-page lock-page" filter-color="black" data-image="/static/assets/img/lock.jpeg">
        <!--   you can change the color of the filter page using: data-color="blue | green | orange | red | purple" -->
        <div class="content">
            <form method="post">
                {{.xsrfdata}}
                <div class="card card-profile card-hidden">
                    <div class="card-avatar">
                        <a href="javascript:void(0)">
                            <img class="avatar" src="/static/assets/img/faces/avatar.jpg" alt="...">
                        </a>
                    </div>
                    <div class="card-content">
                        <h4 class="card-title">Cloud 夜</h4>
                        <div class="form-group label-floating">
                            <label class="control-label">Enter Password</label>
                            <input class="form-control" id="password" name="password" type="password">
                        </div>
                    </div>
                    <div class="card-footer">
                        <button type="submit" class="btn btn-rose btn-round">Unlock</button>
                    </div>
                </div>
            </form>
        </div>
        <footer class="footer">
            <div class="container">
                <nav class="pull-left">
                    <ul>
                        <li>
                            <a href="/">
                                Home
                            </a>
                        </li>
                        <li>
                            <a href="https://github.com/aiyeyun">
                                GitHub
                            </a>
                        </li>
                        <li>
                            <a href="http://blog.aiyeyun.com">
                                Blog
                            </a>
                        </li>
                    </ul>
                </nav>
                <p class="copyright pull-right">
                    &copy;
                    <script>
                        document.write(new Date().getFullYear())
                    </script>
                    <a href="http://www.aiyeyun.com">Cloud 夜</a>,  {{.copyright}}
                </p>
            </div>
        </footer>
    </div>
</div>
</body>
<div class="fixed-plugin">
    <div class="dropdown show-dropdown">
        <a href="#" data-toggle="dropdown">
            <i class="fa fa-cog fa-2x"> </i>
        </a>
        <ul class="dropdown-menu">
            <li class="header-title">Background Style</li>
            <li class="adjustments-line">
                <a href="javascript:void(0)" class="switch-trigger">
                    <p>Background Image</p>
                    <div class="togglebutton switch-sidebar-image">
                        <label>
                            <input type="checkbox" checked="">
                        </label>
                    </div>
                    <div class="clearfix"></div>
                </a>
            </li>
            <li class="adjustments-line">
                <a href="javascript:void(0)" class="switch-trigger active-color">
                    <p>Filters</p>
                    <div class="badge-colors pull-right">
                        <span class="badge filter active" data-color="black"></span>
                        <span class="badge filter badge-blue" data-color="blue"></span>
                        <span class="badge filter badge-green" data-color="green"></span>
                        <span class="badge filter badge-orange" data-color="orange"></span>
                        <span class="badge filter badge-red" data-color="red"></span>
                        <span class="badge filter badge-purple" data-color="purple"></span>
                        <span class="badge filter badge-rose" data-color="rose"></span>
                    </div>
                    <div class="clearfix"></div>
                </a>
            </li>
            <li class="header-title">Background Images</li>
            <li class="active">
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-1.jpg" data-src="/static/assets/img/login.jpeg" alt="" />
                </a>
            </li>
            <li>
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-2.jpg" data-src="/static/assets/img/lock.jpeg" alt="" />
                </a>
            </li>
            <li>
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-3.jpg" data-src="/static/assets/img/header-doc.jpeg" alt="" />
                </a>
            </li>
            <li>
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-4.jpg" data-src="/static/assets/img/bg-pricing.jpeg" alt="" />
                </a>
            </li>
        </ul>
    </div>
</div>
</body>
<!--   Core JS Files   -->
<script src="/static/assets/js/jquery-3.1.1.min.js" type="text/javascript"></script>
<script src="/static/assets/js/jquery-ui.min.js" type="text/javascript"></script>
<script src="/static/assets/js/bootstrap.min.js" type="text/javascript"></script>
<script src="/static/assets/js/material.min.js" type="text/javascript"></script>
<script src="/static/assets/js/perfect-scrollbar.jquery.min.js" type="text/javascript"></script>
<!-- Forms Validations Plugin -->
<script src="/static/assets/js/jquery.validate.min.js"></script>
<!--  Plugin for Date Time Picker and Full Calendar Plugin-->
<script src="/static/assets/js/moment.min.js"></script>
<!--  Charts Plugin -->
<script src="/static/assets/js/chartist.min.js"></script>
<!--  Plugin for the Wizard -->
<script src="/static/assets/js/jquery.bootstrap-wizard.js"></script>
<!--  Notifications Plugin    -->
<script src="/static/assets/js/bootstrap-notify.js"></script>
<!--   Sharrre Library    -->
<script src="/static/assets/js/jquery.sharrre.js"></script>
<!-- DateTimePicker Plugin -->
<script src="/static/assets/js/bootstrap-datetimepicker.js"></script>
<!-- Vector Map plugin -->
<script src="/static/assets/js/jquery-jvectormap.js"></script>
<!-- Sliders Plugin -->
<script src="/static/assets/js/nouislider.min.js"></script>
<!--  Google Maps Plugin    -->
<!--<script src="https://maps.googleapis.com/maps/api/js"></script>-->
<!-- Select Plugin -->
<script src="/static/assets/js/jquery.select-bootstrap.js"></script>
<!--  DataTables.net Plugin    -->
<script src="/static/assets/js/jquery.datatables.js"></script>
<!-- Sweet Alert 2 plugin -->
<script src="/static/assets/js/sweetalert2.js"></script>
<!--	Plugin for Fileupload, full documentation here: http://www.jasny.net/bootstrap/javascript/#fileinput -->
<script src="/static/assets/js/jasny-bootstrap.min.js"></script>
<!--  Full Calendar Plugin    -->
<script src="/static/assets/js/fullcalendar.min.js"></script>
<!-- TagsInput Plugin -->
<script src="/static/assets/js/jquery.tagsinput.js"></script>
<!-- Material Dashboard javascript methods -->
<script src="/static/assets/js/material-dashboard.js"></script>
<!-- Material Dashboard DEMO methods, don't include it in your project! -->
<script src="/static/assets/js/demo.js"></script>
<script src="/static/assets/js/main.js"></script>

<script type="text/javascript">
    $().ready(function() {
        demo.checkFullPageBackgroundImage();

        setTimeout(function() {
            // after 1000 ms we add the class animated to the login/register card
            $('.card').removeClass('card-hidden');
        }, 700)


        $("form").submit(function(e){
            var password = $("#password").val()
            if(!password){
                main.notifyShow("请输入密码", "warning");
                return false;
            }
        });

        {{if .flash.error }}
            main.notifyShow({{ .flash.error }}, "danger")
        {{end}}
    });
</script>

</html>