{{ define "footer" }}
    <!-- jQuery -->
    <script src="{{.URLPathPrefix}}/admin//vendor/jquery/jquery.min.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="{{.URLPathPrefix}}/admin//vendor/bootstrap/js/bootstrap.min.js"></script>

    <!-- Metis Menu Plugin JavaScript -->
    <script src="{{.URLPathPrefix}}/admin//vendor/metisMenu/metisMenu.min.js"></script>

    <!-- Morris Charts JavaScript -->
    <script src="{{.URLPathPrefix}}/admin//vendor/raphael/raphael.min.js"></script>
    <script src="{{.URLPathPrefix}}/admin//vendor/morrisjs/morris.min.js"></script>
    <script src="{{.URLPathPrefix}}/admin//data/morris-data.js"></script>

    <!-- Custom Theme JavaScript -->
    <script src="{{.URLPathPrefix}}/admin//dist/js/sb-admin-2.js"></script>

</body>

</html>
{{ end }}