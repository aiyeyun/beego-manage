<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <link rel="apple-touch-icon" sizes="76x76" href="/static/assets/img/apple-icon.png" />
    <link rel="icon" type="image/png" href="/static/assets/img/favicon.png" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <title>控制台</title>
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
    <link rel="stylesheet" type="text/css" href="/static/assets/css/icons.css" />

    {{.css}}

</head>
<body>

<div class="wrapper">
    <div class="sidebar" data-active-color="rose" data-background-color="black" data-image="/static/assets/img/sidebar-1.jpg">
        <!--
    Tip 1: You can change the color of active element of the sidebar using: data-active-color="purple | blue | green | orange | red | rose"
    Tip 2: you can also add an image using data-image tag
    Tip 3: you can change the color of the sidebar with data-background-color="white | black"
-->
        <div class="logo">
            <a href="http://www.creative-tim.com" class="simple-text">
                夜云
            </a>
        </div>
        <div class="logo logo-mini">
            <a href="http://www.creative-tim.com" class="simple-text">
                云
            </a>
        </div>
        <div class="sidebar-wrapper">
            <div class="user">
                <div class="photo">
                    <img src="/static/assets/img/faces/avatar.jpg" />
                </div>
                <div class="info">
                    <a data-toggle="collapse" href="#collapseExample" class="collapsed">
                        Tania Andrew
                        <b class="caret"></b>
                    </a>
                    <div class="collapse" id="collapseExample">
                        <ul class="nav">
                            <li>
                                <a href="#">My Profile</a>
                            </li>
                            <li>
                                <a href="#">Edit Profile</a>
                            </li>
                            <li>
                                <a href="#">Settings</a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <ul class="nav">
                <li class="active">
                    <a href="./dashboard.html">
                        <i class="material-icons">dashboard</i>
                        <p>Dashboard</p>
                    </a>
                </li>
                <li>
                    <a href="https://material.io/icons/" target="_blank">
                        <i class="material-icons">crop_free</i>
                        <p>Icons</p>
                    </a>
                </li>
                <li>
                    <a data-toggle="collapse" href="#memenumu">
                        <i class="material-icons">menu</i>
                        <p>菜单栏
                            <b class="caret"></b>
                        </p>
                    </a>
                    <div class="collapse" id="menu">
                        <ul class="nav">
                            <li>
                                <a href="http://www.baidu.com">菜单栏列表</a>
                            </li>
                            <li>
                                <a href="/menu/form">添加菜单栏</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a data-toggle="collapse" href="#pagesExamples">
                        <i class="material-icons">image</i>
                        <p>Pages
                            <b class="caret"></b>
                        </p>
                    </a>
                    <div class="collapse" id="pagesExamples">
                        <ul class="nav">
                            <li>
                                <a href="./pages/pricing.html">Pricing</a>
                            </li>
                            <li>
                                <a href="./pages/timeline.html">Timeline</a>
                            </li>
                            <li>
                                <a href="./pages/login.html">Login Page</a>
                            </li>
                            <li>
                                <a href="./pages/register.html">Register Page</a>
                            </li>
                            <li>
                                <a href="./pages/lock.html">Lock Screen Page</a>
                            </li>
                            <li>
                                <a href="./pages/user.html">User Profile</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a data-toggle="collapse" href="#componentsExamples">
                        <i class="material-icons">apps</i>
                        <p>Components
                            <b class="caret"></b>
                        </p>
                    </a>
                    <div class="collapse" id="componentsExamples">
                        <ul class="nav">
                            <li>
                                <a href="./components/buttons.html">Buttons</a>
                            </li>
                            <li>
                                <a href="./components/grid.html">Grid System</a>
                            </li>
                            <li>
                                <a href="./components/panels.html">Panels</a>
                            </li>
                            <li>
                                <a href="./components/sweet-alert.html">Sweet Alert</a>
                            </li>
                            <li>
                                <a href="./components/notifications.html">Notifications</a>
                            </li>
                            <li>
                                <a href="./components/icons.html">Icons</a>
                            </li>
                            <li>
                                <a href="./components/typography.html">Typography</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a data-toggle="collapse" href="#formsExamples">
                        <i class="material-icons">content_paste</i>
                        <p>Forms
                            <b class="caret"></b>
                        </p>
                    </a>
                    <div class="collapse" id="formsExamples">
                        <ul class="nav">
                            <li>
                                <a href="./forms/regular.html">Regular Forms</a>
                            </li>
                            <li>
                                <a href="./forms/extended.html">Extended Forms</a>
                            </li>
                            <li>
                                <a href="./forms/validation.html">Validation Forms</a>
                            </li>
                            <li>
                                <a href="./forms/wizard.html">Wizard</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a data-toggle="collapse" href="#tablesExamples">
                        <i class="material-icons">grid_on</i>
                        <p>Tables
                            <b class="caret"></b>
                        </p>
                    </a>
                    <div class="collapse" id="tablesExamples">
                        <ul class="nav">
                            <li>
                                <a href="./tables/regular.html">Regular Tables</a>
                            </li>
                            <li>
                                <a href="./tables/extended.html">Extended Tables</a>
                            </li>
                            <li>
                                <a href="./tables/datatables.net.html">DataTables.net</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a data-toggle="collapse" href="#mapsExamples">
                        <i class="material-icons">place</i>
                        <p>Maps
                            <b class="caret"></b>
                        </p>
                    </a>
                    <div class="collapse" id="mapsExamples">
                        <ul class="nav">
                            <li>
                                <a href="./maps/google.html">Google Maps</a>
                            </li>
                            <li>
                                <a href="./maps/fullscreen.html">Full Screen Map</a>
                            </li>
                            <li>
                                <a href="./maps/vector.html">Vector Map</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a href="./widgets.html">
                        <i class="material-icons">widgets</i>
                        <p>Widgets</p>
                    </a>
                </li>
                <li>
                    <a href="./charts.html">
                        <i class="material-icons">timeline</i>
                        <p>Charts</p>
                    </a>
                </li>
                <li>
                    <a href="./calendar.html">
                        <i class="material-icons">date_range</i>
                        <p>Calendar</p>
                    </a>
                </li>
            </ul>
        </div>
    </div>
    <div class="main-panel">
        <nav class="navbar navbar-transparent navbar-absolute">
            <div class="container-fluid">
                <div class="navbar-minimize">
                    <button id="minimizeSidebar" class="btn btn-round btn-white btn-fill btn-just-icon">
                        <i class="material-icons visible-on-sidebar-regular">more_vert</i>
                        <i class="material-icons visible-on-sidebar-mini">view_list</i>
                    </button>
                </div>
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="#"> Dashboard </a>

                    {{.navbar}}

                </div>
                <div class="collapse navbar-collapse">
                    <ul class="nav navbar-nav navbar-right">
                        <li>
                            <a href="#pablo" class="dropdown-toggle" data-toggle="dropdown">
                                <i class="material-icons">dashboard</i>
                                <p class="hidden-lg hidden-md">Dashboard</p>
                            </a>
                        </li>
                        <li class="dropdown">
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                <i class="material-icons">notifications</i>
                                <span class="notification">5</span>
                                <p class="hidden-lg hidden-md">
                                    Notifications
                                    <b class="caret"></b>
                                </p>
                            </a>
                            <ul class="dropdown-menu">
                                <li>
                                    <a href="#">Mike John responded to your email</a>
                                </li>
                                <li>
                                    <a href="#">You have 5 new tasks</a>
                                </li>
                                <li>
                                    <a href="#">You're now friend with Andrew</a>
                                </li>
                                <li>
                                    <a href="#">Another Notification</a>
                                </li>
                                <li>
                                    <a href="#">Another One</a>
                                </li>
                            </ul>
                        </li>
                        <li>
                            <a href="#pablo" class="dropdown-toggle" data-toggle="dropdown">
                                <i class="material-icons">person</i>
                                <p class="hidden-lg hidden-md">Profile</p>
                            </a>
                        </li>
                        <li class="logout">
                            <a href="/logout" class="itemBox" id="logout" data-toggle="dropdown">
                                <i class="material-icons">power_settings_new</i>
                            </a>
                        </li>
                        <li class="separator hidden-lg hidden-md"></li>
                    </ul>
                    <form class="navbar-form navbar-right" role="search">
                        <div class="form-group form-search is-empty">
                            <input type="text" class="form-control" placeholder="Search">
                            <span class="material-input"></span>
                        </div>
                        <button type="submit" class="btn btn-white btn-round btn-just-icon">
                            <i class="material-icons">search</i>
                            <div class="ripple-container"></div>
                        </button>
                    </form>
                </div>
            </div>
        </nav>
        <div class="content">
            {{.LayoutContent}}
        </div>
        <footer class="footer">
            <div class="container-fluid">
                <nav class="pull-left">
                    <ul>
                        <li>
                            <a href="/console">
                                Home
                            </a>
                        </li>
                        <li>
                            <a href="http://www.aiyeyun.com">
                                Company
                            </a>
                        </li>
                        <li>
                            <a href="https://github.com/aiyeyun">
                                Github
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

<div class="fixed-plugin">
    <div class="dropdown show-dropdown">
        <a href="#" data-toggle="dropdown">
            <i class="fa fa-cog fa-2x"> </i>
        </a>
        <ul class="dropdown-menu">
            <!-- <li class="header-title"> Sidebar Filters</li> -->
            <li class="header-title"> 侧边栏过滤器</li>
            <li class="adjustments-line">
                <a href="javascript:void(0)" class="switch-trigger active-color">
                    <div class="badge-colors text-center">
                        <span class="badge filter badge-purple" data-color="purple"></span>
                        <span class="badge filter badge-blue" data-color="blue"></span>
                        <span class="badge filter badge-green" data-color="green"></span>
                        <span class="badge filter badge-orange" data-color="orange"></span>
                        <span class="badge filter badge-red" data-color="red"></span>
                        <span class="badge filter badge-rose active" data-color="rose"></span>
                    </div>
                    <div class="clearfix"></div>
                </a>
            </li>
            <!-- <li class="header-title">Sidebar Background</li> -->
            <li class="header-title">侧边栏背景</li>
            <li class="adjustments-line">
                <a href="javascript:void(0)" class="switch-trigger background-color">
                    <div class="text-center">
                        <span class="badge filter badge-white" data-color="white"></span>
                        <span class="badge filter badge-black active" data-color="black"></span>
                    </div>
                    <div class="clearfix"></div>
                </a>
            </li>
            <li class="adjustments-line">
                <a href="javascript:void(0)" class="switch-trigger">
                    <!-- <p>Sidebar Mini</p> -->
                    <p>迷你侧边栏</p>
                    <div class="togglebutton switch-sidebar-mini">
                        <label>
                            <input type="checkbox" unchecked="">
                        </label>
                    </div>
                    <div class="clearfix"></div>
                </a>
            </li>
            <li class="adjustments-line">
                <a href="javascript:void(0)" class="switch-trigger">
                    <!-- <p>Sidebar Image</p> -->
                    <p>侧边栏的图像</p>
                    <div class="togglebutton switch-sidebar-image">
                        <label>
                            <input type="checkbox" checked="">
                        </label>
                    </div>
                    <div class="clearfix"></div>
                </a>
            </li>
            <!-- <li class="header-title">Images</li> -->
            <li class="header-title">图像</li>
            <li class="active">
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-1.jpg" alt="" />
                </a>
            </li>
            <li>
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-2.jpg" alt="" />
                </a>
            </li>
            <li>
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-3.jpg" alt="" />
                </a>
            </li>
            <li>
                <a class="img-holder switch-trigger" href="javascript:void(0)">
                    <img src="/static/assets/img/sidebar-4.jpg" alt="" />
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
<script type="text/javascript">
    $(document).ready(function() {
        $(".click-li").click(function () {
            location.href= $(this).find('a').attr('href')
        })
    });
</script>

{{.js}}

</html>