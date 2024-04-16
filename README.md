> # ⚠️ This is deprecated!! [`osv-scanner`](https://google.github.io/osv-scanner) by Google does the work way better than this tool 

![](img/nodeps_banner.png)

Nodeps (**No** **Dep**endencie**s**) will check which libraries in `package.json` are up to date and which ones are using an outdated version which might have vulnerabilities.

## Usage

```
$ nodeps -pkg package.json

     |  NAME        |  USED VERSION  |  LAST VERSION  |  REPOSITORY
-----+--------------+----------------+----------------+---------------------------------+
  ✅ | lib1         | ^4.0.0         | 4.0.0          | https://github.com/user/lib1
  ✅ | lib2         | ^5.0.5         | 5.0.5          | https://github.com/user/lib2
  ⚠️  | lib3         | ^0.1.0         | 0.21.1         | https://github.com/user/lib3
  ✅ | lib4         | ^1.5.0         | 1.5.0          | https://github.com/user/lib4

```

