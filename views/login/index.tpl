<!doctype html>
<head>
    <meta charset="utf-8">
    <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;" name="viewport" />
    <meta name="_xsrf" content="{{.xsrf_token}}" />

    <link href="/static/assets/css/bootstrap.min.css" rel="stylesheet" />
    <link href="/static/assets/css/font-awesome.min.css" rel="stylesheet">
    <link href="/static/assets/css/material-dashboard.css" rel="stylesheet" />
    <link href="/static/assets/css/icons.css" rel="stylesheet" />
    <link href="/static/css/login.css" rel="stylesheet">
</head>
<body>
<div id="container">
    <video id="background_video" loop muted></video>
    <div id="video_cover"></div>
    <div id="overlay"></div>
    <section id="main_content">

        <div class="full-page login-page" filter-color="black">
            <!--   you can change the color of the filter page using: data-color="blue | purple | green | orange | red | rose " -->
            <div class="content">
                <div class="container">
                    <div class="row">
                        <div class="col-md-4 col-sm-6 col-md-offset-4 col-sm-offset-3">

                            <div class="card card-login">
                                <div class="card-header text-center" data-background-color="rose">
                                    <h3 class="card-title">Manage Login</h3>
                                    <div class="social-line">
                                        <a href="#btn" class="btn btn-just-icon btn-simple">
                                            <i class="fa fa-github"></i>
                                        </a>
                                        <a href="#pablo" class="btn btn-just-icon btn-simple">
                                            <i class="fa fa-twitter"></i>
                                        </a>
                                        <a href="#eugen" class="btn btn-just-icon btn-simple">
                                            <i class="fa fa-google-plus"></i>
                                        </a>
                                    </div>
                                </div>

                                <p class="category text-center">
                                    Welcome to cloud
                                </p>

                                {{if .flash.error }}
                                    <div class="alert alert-danger">
                                        <span>{{.flash.error}}</span>
                                    </div>
                                {{end}}

                                <form id="login-form" method="post" novalidate="novalidate">
                                    {{.xsrfdata}}
                                    <div class="card-content">
                                        <div class="input-group">
                                            <span class="input-group-addon">
                                                <i class="material-icons">face</i>
                                            </span>
                                            <div class="form-group label-floating is-empty">
                                                <label class="control-label">Username</label>
                                                <input class="form-control" name="username" type="text" required="true" aria-required="true">
                                            </div>
                                        </div>
                                        <div class="input-group">
                                            <span class="input-group-addon">
                                                <i class="material-icons">lock_outline</i>
                                            </span>
                                            <div class="form-group label-floating is-empty">
                                                <label class="control-label">Password</label>
                                                <input class="form-control" name="password" type="password" required="true" aria-required="true">
                                                <span class="material-input"></span>
                                            </div>
                                        </div>
                                    </div>

                                    <div class="footer text-center">
                                        <button type="submit" class="btn btn-rose btn-simple btn-wd btn-lg">Let's go</button>
                                    </div>
                                </form>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </div>

    </section>
</div>
<script src="/static/js/bideo.js"></script>
<script src="/static/js/login.js"></script>

<!--   Core JS Files   -->
<script src="/static/assets/js/jquery-3.1.1.min.js" type="text/javascript"></script>
<script src="/static/assets/js/bootstrap.min.js" type="text/javascript"></script>
<script src="/static/assets/js/material.min.js" type="text/javascript"></script>
<!-- Forms Validations Plugin -->
<script src="/static/assets/js/jquery.validate.min.js"></script>
<script src="/static/assets/js/jquery.tagsinput.js"></script>
<!-- Material Dashboard javascript methods -->
<script src="/static/assets/js/material-dashboard.js"></script>

<!--  Notifications Plugin    -->
<script src="/static/assets/js/bootstrap-notify.js"></script>
<!--   Sharrre Library    -->
<script src="/static/assets/js/jquery.sharrre.js"></script>
<script src="/static/assets/js/main.js"></script>


<script>
    function setFormValidation(id) {
        $(id).validate({
            errorPlacement: function(error, element) {
                $(element).parent('div').addClass('has-error');
            }
        });
    }

    $(document).ready(function() {
        setFormValidation('#login-form');
        {{if .flash.error }}
            main.notifyShow({{ .flash.error }}, "danger")
        {{end}}

    });
</script>
</body>
</html>


