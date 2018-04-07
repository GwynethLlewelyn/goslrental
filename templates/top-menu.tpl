{{ define "top-menu" }}
			<div class="navbar-header">
				<button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
				<a class="navbar-brand" href="{{.URLPathPrefix}}/admin/"><img src="{{.URLPathPrefix}}/lib/images/favicon-32x32.png" height="32" alt="{{.Title}}"></a>
			</div>
			<!-- /.navbar-header -->
			<ul class="nav navbar-top-links navbar-right">
				                <li class="dropdown">
                    <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                        <i class="fa fa-envelope fa-fw"></i> <i class="fa fa-caret-down"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-messages">
                        <li>
                            <a href="#">
                                <div>
                                    <strong>John Smith</strong>
                                    <span class="pull-right text-muted">
                                        <em>Yesterday</em>
                                    </span>
                                </div>
                                <div>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eleifend...</div>
                            </a>
                        </li>
                        <li class="divider"></li>
                        <li>
                            <a href="#">
                                <div>
                                    <strong>John Smith</strong>
                                    <span class="pull-right text-muted">
                                        <em>Yesterday</em>
                                    </span>
                                </div>
                                <div>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eleifend...</div>
                            </a>
                        </li>
                        <li class="divider"></li>
                        <li>
                            <a href="#">
                                <div>
                                    <strong>John Smith</strong>
                                    <span class="pull-right text-muted">
                                        <em>Yesterday</em>
                                    </span>
                                </div>
                                <div>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eleifend...</div>
                            </a>
                        </li>
                        <li class="divider"></li>
                        <li>
                            <a class="text-center" href="#">
                                <strong>Read All Messages</strong>
                                <i class="fa fa-angle-right"></i>
                            </a>
                        </li>
                    </ul>
                    <!-- /.dropdown-messages -->
                </li>
                <!-- /.dropdown -->
				{{ if .SetCookie }}
				 <li id="#username">
				{{ .SetCookie }}
				{{ if .GravatarMenu }}
					 <div style="float:left;" class="gravatar-container">
						  <a href="https://gravatar.com/{{ .GravatarHash }}" title="{{ .SetCookie }}">
							   <img class="avatar avatar-{{ .GravatarSizeMenu }} photo" src="{{ .GravatarMenu }}" srcset="https://secure.gravatar.com/avatar/{{ .GravatarHash }}?s={{ .GravatarTwiceSizeMenu }}&amp;d=mm&amp;r=r 2x" height="{{ .GravatarSizeMenu }}" width="{{ .GravatarSizeMenu }}" alt="{{ .SetCookie }}">
						</a>
					</div> <!-- ./gravatar-container -->
				{{ end }}
				 </li> <!-- ./username -->
				{{ end }}
<!--				<li>
					<a href="{{.URLPathPrefix}}/admin/"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
				</li> -->
				<li class="dropdown">
					<a class="dropdown-toggle" data-toggle="dropdown" href="#">
						<i class="fa fa-database fa-fw"></i> Database <i class="fa fa-caret-down"></i>
					</a>
					<ul class="dropdown-menu">
						  <li>
							  <a href="{{.URLPathPrefix}}/admin/agents/"><i class="fa fa-android fa-fw"></i> Agents (Bots/NPCs)</a>
						  </li>
						  <li>
							  <a href="{{.URLPathPrefix}}/admin/positions/"><i class="fa fa-codepen fa-fw"></i> Positions (Cubes)</a>
						  </li>
						  <li>
							  <a href="{{.URLPathPrefix}}/admin/inventory/"><i class="fa fa-folder-open-o fa-fw"></i> Content/Inventory</a>
						  </li>
						  <li>
							  <a href="{{.URLPathPrefix}}/admin/objects/"><i class="fa fa-cubes fa-fw"></i> Obstacles (Objects)</a>
						  </li>
					</ul>
					<!-- /.dropdown-menu -->
				</li>
				<li class="dropdown">
                    <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                        <i class="fa fa-user fa-fw"></i> <i class="fa fa-caret-down"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-user">
                        <li><a href="{{.URLPathPrefix}}/admin/user-management/"><i class="fa fa-users fa-fw"></i> Manage Users</a>
                        </li>
                        <li><a href="#"><i class="fa fa-gear fa-fw"></i> Settings</a>
                        </li>
                        <li class="divider"></li>
                        <li><a href="{{.URLPathPrefix}}/admin/logout/"><i class="fa fa-sign-out fa-fw"></i> Logout</a>
                        </li>
                    </ul>
                    
                </li>
			</ul>
			<!-- /.navbar-top-links -->		  
{{ end }}