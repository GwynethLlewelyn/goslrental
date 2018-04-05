{{ define "sidebar-left" }}
			<div class="navbar-default sidebar" role="navigation">
				<div class="sidebar-nav navbar-collapse">
					<ul class="nav" id="side-menu">
						<li class="sidebar-search">
							<div class="input-group custom-search-form">
								<input type="text" class="form-control" placeholder="Search...">
								<span class="input-group-btn">
								<button class="btn btn-default" type="button">
									<i class="fa fa-search"></i>
								</button>
							</span>
							</div>
							<!-- /input-group -->
						</li>		 
						<li>
							<a href="{{.URLPathPrefix}}/admin/"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
						</li>
						<li>
							<a href="#"><i class="fa fa-bar-chart-o fa-fw"></i> Charts<span class="fa arrow"></span></a>
							<ul class="nav nav-second-level">
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/flot.html">Flot Charts</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/morris.html">Morris.js Charts</a>
								</li>
							</ul>
							<!-- /.nav-second-level -->
						</li>
						<li>
							<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/tables.html"><i class="fa fa-table fa-fw"></i> Tables</a>
						</li>
						<li>
							<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/forms.html"><i class="fa fa-edit fa-fw"></i> Forms</a>
						</li>
						<li>
							<a href="#"><i class="fa fa-wrench fa-fw"></i> UI Elements<span class="fa arrow"></span></a>
							<ul class="nav nav-second-level">
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/panels-wells.html">Panels and Wells</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/buttons.html">Buttons</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/notifications.html">Notifications</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/typography.html">Typography</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/icons.html"> Icons</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/grid.html">Grid</a>
								</li>
							</ul>
							<!-- /.nav-second-level -->
						</li>
						<li>
							<a href="#"><i class="fa fa-sitemap fa-fw"></i> Multi-Level Dropdown<span class="fa arrow"></span></a>
							<ul class="nav nav-second-level">
								<li>
									<a href="#">Second Level Item</a>
								</li>
								<li>
									<a href="#">Second Level Item</a>
								</li>
								<li>
									<a href="#">Third Level <span class="fa arrow"></span></a>
									<ul class="nav nav-third-level">
										<li>
											<a href="#">Third Level Item</a>
										</li>
										<li>
											<a href="#">Third Level Item</a>
										</li>
										<li>
											<a href="#">Third Level Item</a>
										</li>
										<li>
											<a href="#">Third Level Item</a>
										</li>
									</ul>
									<!-- /.nav-third-level -->
								</li>
							</ul>
							<!-- /.nav-second-level -->
						</li>
						<li>
							<a href="#"><i class="fa fa-files-o fa-fw"></i> Sample Pages<span class="fa arrow"></span></a>
							<ul class="nav nav-second-level">
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/blank.html">Blank Page</a>
								</li>
								<li>
									<a href="{{.URLPathPrefix}}/lib/startbootstrap-sb-admin-2/pages/login.html">Login Page</a>
								</li>
							</ul>
							<!-- /.nav-second-level -->
						</li>
					</ul> <!-- /.side-menu -->
				</div>
				<!-- /.sidebar-collapse -->
			</div>
			<!-- /.navbar-static-side -->
		
{{ end }}