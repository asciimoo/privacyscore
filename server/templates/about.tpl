{{ define "content" }}
<h2>About PrivacyScore</h2>
<div class="row">
    <div class="column column-80 column-offset-10">
        <p>PrivacyScore is a static checker tool that reveals privacy issues.<br />
        Using PrivacyScore you can identify elements on websites which might violate your privacy.<br />
        The check isn't comprehensive yet, but it can be a compass in the digital world.<br />
        If you have suggestion, idea or question, you can open an issue on <a href="https://github.com/asciimoo/privacyscore">GitHub</a>.</p>
    </div>
</div>

<h3>Penalties</h3>

<div class="penalty_info" id="p_external_link">
    <h4 class="left">Send HTTP referrer to foreign host</h4>
    <h5 class="right">1 point per host</h5>
    <p>Referrer HTTP header identifies the address of the webpage that linked to the resource being requested</p>
    <p>Solution is the "referrer" HTML meta tag:<pre> &lt;meta name="referrer" content="none"/&gt;</pre></p>
</div>

<div class="penalty_info" id="p_external_resource">
    <h4 class="left">Load external resource</h4>
    <h5 class="right">5 points per host</h5>
    <p>Leaks tons of information to remote service</p>
    <p>Solution: serve your resources from the same host</p>
</div>

<div class="penalty_info" id="p_no_https">
    <h4 class="left">No HTTPS</h4>
    <h5 class="right">6 points</h5>
    <p>Makes possible to sniff the traffic between endpoints</p>
    <p>Solution: use HTTPS - see <a href="https://letsencrypt.org/">Let's Encrypt</a></p>
</div>

<div class="penalty_info" id="p_js">
    <h4 class="left">JavaScript required</h4>
    <h5 class="right">7 points</h5>
    <p>JavaScript is a deep swamp of user tracking</p>
    <p>Solution: serve your page without JavaScript</p>
</div>

<div class="penalty_info" id="p_cookie">
    <h4 class="left">Automatically set cookies</h4>
    <h5 class="right">3 points</h5>
    <p>Cookie can be used for user session tracking</p>
    <p>Solution: do not send cookies</p>
</div>

<div class="penalty_info" id="p_no_secure_header">
    <h4 class="left">Missing secure HTTP headers</h4>
    <h5 class="right">3 points per header</h5>
    <p>Prefered values:
        <pre>
 X-Frame-Options: DENY # SAMEORIGIN also accepted
 X-Xss-Protection: 1; mode=block
 X-Content-Type-Options: nosniff
 Strict-Transport-Security: max-age=31536000 # only on HTTPS</pre>
    </p>
</div>

<div class="penalty_info" id="p_iframe">
    <h4 class="left">External content in iFrame</h4>
    <h5 class="right">5 points per host</h5>
    <p>Leaks tons of information to remote service</p>
    <p>Solution: serve your page without external iFrames</p>
</div>

<div class="penalty_info" id="p_http_link">
    <h4 class="left">HTTP link</h4>
    <h5 class="right">1 point per host</h5>
    <p>Shepherds users to use unencrypted services</p>
    <p>Solution: use HTTPS links</p>
</div>
{{ end }}
