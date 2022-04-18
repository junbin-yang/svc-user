
%define user_group dvs
%define installpath /usr/local/svc-user
%define logpath /var/log/svc-user
%define sharepath /usr/local/svc-user/docs
%define webpath /usr/local/svc-user/web

Name:           svc-user
License:        GPL
Group:          sunlight
Url:            www.sunlight-tech.com
Version:        _VERSION_
Release:        _RELEASE_
Summary:        User Managed System
Source:         %{name}.%{version}.%{release}.tar.gz
Source1:        %{name}.service
Source2:        %{name}.yml
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%if 0%{?fedora} >= 15 || 0%{?rhel} >=7 || 0%{?suse_version} >= 1210
%global use_systemd 1
%else
%global use_systemd 0
%endif

%if 0%{?use_systemd}
%{?systemd_requires}
%if 0%{?suse_version}
BuildRequires:   systemd-rpm-macros
%else
BuildRequires:   systemd
%endif
%endif

Conflicts:      %{name} <= %{version}

%description

%prep
%setup -c

%install
install -d $RPM_BUILD_ROOT%{installpath}
install -d $RPM_BUILD_ROOT%{sharepath}
install -d $RPM_BUILD_ROOT%{webpath}
cp -a %{name}.%{version}.%{release}/%{name} $RPM_BUILD_ROOT%{installpath}/
cp -a %{name}.%{version}.%{release}/docs/* $RPM_BUILD_ROOT%{sharepath}/
cp -a %{name}.%{version}.%{release}/web/* $RPM_BUILD_ROOT%{webpath}/
install -D -m 0644 %{SOURCE1} %{buildroot}%{_unitdir}/%{name}.service
install -D -m 0644 %{SOURCE2} $RPM_BUILD_ROOT/etc/%{name}.yml

%post
%if ! 0%{?use_systemd}
/sbin/chkconfig --add %{name} > /dev/null 2>&1 || :
%else
%if 0%{?suse_version}
%service_add_post %{name}.service
%else
%systemd_post %{name}.service
%endif
%endif
mkdir -p %{logpath}/%{name}
chown -R %{user_group}:%{user_group} %{logpath}

%postun
%if 0%{?use_systemd}
%if 0%{?suse_version}
%service_del_postun %{name}.service
%else
%systemd_postun_with_restart %{name}.service
%endif
%endif

%pre 
%if 0%{?use_systemd} && 0%{?suse_version}
%service_add_pre %{name}.service
%endif
%{_sbindir}/groupadd -r %{user_group} &>/dev/null ||:
%{_sbindir}/useradd -g %{user_group} -s /bin/false -r -c "user for %{user_group}" %{user_group} &>/dev/null ||:

%preun
%if 0%{?suse_version}
%service_del_preun %{name}.service
%else
%systemd_preun %{name}.service
%endif

%clean
[ "$RPM_BUILD_ROOT" != "/" ] && rm -rf $RPM_BUILD_ROOT
rm -rf $RPM_BUILD_DIR/%{name}-%{version}

%files
%defattr(-, dvs, dvs)
%config(noreplace) /etc/%{name}.yml
%{_unitdir}/%{name}.service
%{installpath}
%{sharepath}
%{webpath}