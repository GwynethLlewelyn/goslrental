{{ define "header" }}<!DOCTYPE html>
<html lang="en">
<head>

	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="description" content="{{.Title}}">
	<meta name="author" content="Gwyneth Llewelyn">

	<title>{{.Title}}</title>

	<!-- Google Web Fonts -->
	<link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Cantarell|Cardo">
	
	<!-- Bootstrap Core CSS -->
	<link href="/lib/startbootstrap-sb-admin-2/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">
	<!--<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">-->


	<!-- MetisMenu CSS -->
	<link href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/vendor/metisMenu/metisMenu.min.css" rel="stylesheet">

	<!-- Custom CSS -->
	<link href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/dist/css/sb-admin-2.css" rel="stylesheet">

	<!-- Morris Charts CSS -->
	<link href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/vendor/morrisjs/morris.css" rel="stylesheet">

	<!-- Custom Fonts -->
	<link href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
	<!--<script defer src="https://use.fontawesome.com/releases/v5.0.9/js/all.js" integrity="sha384-8iPTk2s/jMVj81dnzb/iFR2sdA7u06vHJyyLlAd4snFpCl/SnyUjRrbdJsw1pGIl" crossorigin="anonymous"></script>-->
	
	{{ if .Gravatar }}
	<!-- I have no idea if this is really needed! -->
	<link rel="stylesheet" href="https://secure.gravatar.com/css/services.css" type="text/css">
	<link rel="stylesheet" href="{{.URLPathPrefix}}/lib/gravatar-profile.css" type="text/css">
	{{ end }}
	
	{{ if .LSL }}
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.6.0/themes/prism.min.css" type="text/css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.6.0/prism.min.js"></script>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.6.0/plugins/toolbar/prism-toolbar.min.css" type="text/css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.6.0/plugins/toolbar/prism-toolbar.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.6.0/plugins/copy-to-clipboard/prism-copy-to-clipboard.min.js"></script>
	{{ end }}

	<!-- Our own overrides -->
	<link href="{{.URLPathPrefix}}/lib/goslrental.css" rel="stylesheet" type="text/css">
	

	<!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
	<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
	<!--[if lt IE 9]>
		<script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
		<script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
	<![endif]-->
	<link rel="apple-touch-icon" sizes="180x180" href="/lib/images/apple-touch-icon.png">
	<link rel="icon" type="image/png" sizes="32x32" href="/lib/images//favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="194x194" href="/lib/images/favicon-194x194.png">
	<link rel="icon" type="image/png" sizes="192x192" href="/lib/images/android-chrome-192x192.png">
	<link rel="icon" type="image/png" sizes="16x16" href="/lib/images/favicon-16x16.png">
	<link rel="manifest" href="/lib/images//site.webmanifest">
	<link rel="mask-icon" href="/lib/images/safari-pinned-tab.svg" color="#5bbad5">
	<meta name="apple-mobile-web-app-title" content="Go SL Rental">
	<meta name="application-name" content="Go SL Rental">
	<meta name="msapplication-TileColor" content="#2b5797">
	<meta name="msapplication-TileImage" content="/lib/images/mstile-144x144.png">
	<meta name="theme-color" content="#ffffff">
</head>

<body>
{{ if .Gravatar }}
<!-- Gravatar Hovercards are sneaky, they add their own CSS at the top of the header! -->
<style>
.gcard {
	z-index: 1000;
}
.emptyPlaceholder {
	z-index: 1000;
}
</style>
{{ end }}
<span id="URLPathPrefix" hidden>{{.URLPathPrefix}}</span>
{{ end }}