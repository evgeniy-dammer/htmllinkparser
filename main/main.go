package main

import (
	"fmt"
	"strings"

	link "github.com/evgeniy-dammer/htmllinkparser"
)

var exampleHtml1 = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page
		<span> some span </span>	
	</a>
</body>
</html>
`

var exampleHtml2 = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`

var exampleHtml3 = `
<!DOCTYPE html>
<!--[if lt IE 7]> <html class="ie ie6 lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="ie ie7 lt-ie9 lt-ie8"        lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="ie ie8 lt-ie9"               lang="en"> <![endif]-->
<!--[if IE 9]>    <html class="ie ie9"                      lang="en"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en" class="no-ie">
<!--<![endif]-->

<head>
  <title> hhf ejvok modfkm dkm f oeofivm fo ,f fvdv</title>
</head>

<body>
  <section class="header-section">
    <div class="jumbo-content">
      <div class="pull-right login-section">
        Already have an account?
        <a href="#" class="btn btn-login">Login <i class="fa fa-sign-in" aria-hidden="true"></i></a>
      </div>
      <center>
        <img src="https://blabla.com/img/logo.png" style="max-width: 85%; z-index: 3;">
        <h1>dbciuwe dvi fojndfovijedf dfv</h1>
        <br/>
        <form action="/do-stuff" method="post">
          <div class="input-group">
            <input type="email" id="drip-email" name="fields[email]" class="btn-input" placeholder="Email Address" required>
            <button class="btn btn-success btn-lg" type="submit">Sign me up!</button>
            <a href="/lost">Lost? Need help?</a>
          </div>
        </form>
        <p class="disclaimer disclaimer-box">jhbvd kjfn ldkfmk mpdfbdpf*&Hcifv;fvn dofvoaiddfnv kfmvldkvlkmfvdfv.</p>
      </center>
    </div>
  </section>
  <section class="footer-section">
    <div class="row">
      <div class="col-md-6 col-md-offset-1 vcenter">
        <div class="quote">
          ucbe7yeu 8uehrvuj 9fijvifv kj ijfviejplc dokvmdfkmvklmd dofkvmdokfvdokfvnoi efovkrvok
        </div>
      </div>
      <div class="col-md-4 col-md-offset-0 vcenter">
        <center>
          <img src="https://blabla.com/img/lifting.gif" style="width: 80%">
          <br/>
          <br/>
        </center>
      </div>
    </div>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <center>
          <p class="disclaimer">
            kvjnsdvjn flvkdmfvlkm sdvkdlvk (<a href="https://blabla.com/bla">@blablabla</a>), cbshdv dv ddckec skcl (fivvue7 kme0 3vm!), nvievunrivj fijnirdfjn  fdfjvnifjn dfvnv.
          </p>
        </center>
      </div>
    </div>
  </section>
</body>
</html>
`

var exampleHtml4 = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml1)

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}
