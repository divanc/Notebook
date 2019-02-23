# Git-flow Aliases

Full table is available [here](https://danielkummer.github.io/git-flow-cheatsheet/img/git-flow-commands.png).

![cheatsheet](https://www.google.com/url?sa=i&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwi78pHSu9LgAhXNp4sKHe_WAcsQjRx6BAgBEAU&url=https%3A%2F%2Fdanielkummer.github.io%2Fgit-flow-cheatsheet%2F&psig=AOvVaw3VIk6-U08CRAriUORZAX_Q&ust=1551032277694674)

<table>
<thead>
<tr>
<th>Alias</th>
<th>Command</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>gfl</code></td>
<td><code>git flow</code></td>
<td>Git-Flow command</td>
</tr>
<tr>
<td><code>gfli</code></td>
<td><code>git flow init</code></td>
<td>Initialize git-flow repository</td>
</tr>
<tr>
<td><code>gcd</code></td>
<td><code>git checkout develop</code></td>
<td>Check out develop branch</td>
</tr>
<tr>
<td><code>gch</code></td>
<td><code>git checkout hotfix</code></td>
<td>Check out hotfix branch</td>
</tr>
<tr>
<td><code>gcr</code></td>
<td><code>git checkout release</code></td>
<td>Check out release branch</td>
</tr>
<tr>
<td><code>gflf</code></td>
<td><code>git flow feature</code></td>
<td>List existing feature branches</td>
</tr>
<tr>
<td><code>gflh</code></td>
<td><code>git flow hotfix</code></td>
<td>List existing hotfix branches</td>
</tr>
<tr>
<td><code>gflr</code></td>
<td><code>git flow release</code></td>
<td>List existing release branches</td>
</tr>
<tr>
<td><code>gflfs</code></td>
<td><code>git flow feature start</code></td>
<td>Start a new feature: <code>gflfs &lt;name&gt;</code></td>
</tr>
<tr>
<td><code>gflhs</code></td>
<td><code>git flow hotfix start</code></td>
<td>Start a new hotfix: <code>gflhs &lt;version&gt;</code></td>
</tr>
<tr>
<td><code>gflrs</code></td>
<td><code>git flow release start</code></td>
<td>Start a new release: <code>gflrs &lt;version&gt;</code></td>
</tr>
<tr>
<td><code>gflff</code></td>
<td><code>git flow feature finish</code></td>
<td>Finish feature: <code>gflff &lt;name&gt;</code></td>
</tr>
<tr>
<td><code>gflfp</code></td>
<td><code>git flow feature publish</code></td>
<td>Publish feature: <code>gflfp &lt;name&gt;</code></td>
</tr>
<tr>
<td><code>gflhf</code></td>
<td><code>git flow hotfix finish</code></td>
<td>Finish hotfix: <code>gflhf &lt;version&gt;</code></td>
</tr>
<tr>
<td><code>gflrf</code></td>
<td><code>git flow release finish</code></td>
<td>Finish release: <code>gflrf &lt;version&gt;</code></td>
</tr>
</tbody>
</table>

# Git zsh aliases

Full table is available [here](https://github.com/robbyrussell/oh-my-zsh/wiki/Plugin:git).

<table>
<thead>
<tr>
<th align="left">Alias</th>
<th align="left">Command</th>
</tr>
</thead>
<tbody>
<tr>
<td align="left">g</td>
<td align="left">git</td>
</tr>
<tr>
<td align="left">gl</td>
<td align="left">git pull</td>
</tr>
<tr>
<td align="left">glola</td>
<td align="left">git log --graph --pretty = format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)&lt;%an&gt;%Creset' --abbrev-commit --all</td>
</tr>
<tr>
<td align="left">ga</td>
<td align="left">git add</td>
</tr>
<tr>
<td align="left">gaa</td>
<td align="left">git add --all</td>
</tr>
<tr>
<td align="left">gapa</td>
<td align="left">git add --patch</td>
</tr>
<tr>
<td align="left">gau</td>
<td align="left">git add --update</td>
</tr>
<tr>
<td align="left">gb</td>
<td align="left">git branch</td>
</tr>
<tr>
<td align="left">gba</td>
<td align="left">git branch -a</td>
</tr>
<tr>
<td align="left">gm</td>
<td align="left">git merge</td>
</tr>
<tr>
<td align="left">gmom</td>
<td align="left">git merge origin/master</td>
</tr>
<tr>
<td align="left">gcmsg</td>
<td align="left">git commit -m</td>
</tr>
<tr>
<td align="left">gc</td>
<td align="left">git commit -v</td>
</tr>
<tr>
<td align="left">gc!</td>
<td align="left">git commit -v --amend</td>
</tr>
<tr>
<td align="left">gco</td>
<td align="left">git checkout</td>
</tr>
<tr>
<td align="left">gcm</td>
<td align="left">git checkout master</td>
</tr>
<tr>
<td align="left">gcd</td>
<td align="left">git checkout develop</td>
</tr>
<tr>
<td align="left">gcb</td>
<td align="left">git checkout -b</td>
</tr>
<tr>
<td align="left">gp</td>
<td align="left">git push</td>
</tr>
<tr>
<td align="left">gd</td>
<td align="left">git diff</td>
</tr>
<tr>
<td align="left">gcf</td>
<td align="left">git config --list</td>
</tr>
<tr>
<td align="left">gcl</td>
<td align="left">git clone --recursive</td>
</tr>
<tr>
<td align="left">gf</td>
<td align="left">git fetch</td>
</tr>
<tr>
<td align="left">gclean</td>
<td align="left">git clean -df</td>
</tr>
<tr>
<td align="left">gcount</td>
<td align="left">git shortlog -sn</td>
</tr>
<tr>
<td align="left">gcp</td>
<td align="left">git cherry-pick</td>
</tr>
<tr>
<td align="left">gcpa</td>
<td align="left">git cherry-pick --abort</td>
</tr>
<tr>
<td align="left">gcpc</td>
<td align="left">git cherry-pick --continue</td>
</tr>
<tr>
<td align="left">gcs</td>
<td align="left">git commit -S</td>
</tr>

<tr>
<td align="left">gdca</td>
<td align="left">git diff --cached</td>
</tr>
<tr>
<td align="left">gdt</td>
<td align="left">git diff-tree --no-commit-id --name-only -r</td>
</tr>
<tr>
<td align="left">gdw</td>
<td align="left">git diff --word-diff</td>
</tr>

<tr>
<td align="left">gfa</td>
<td align="left">git fetch --all --prune</td>
</tr>
<tr>
<td align="left">gfo</td>
<td align="left">git fetch origin</td>
</tr>
<tr>
<td align="left">gg</td>
<td align="left">git gui citool</td>
</tr>
<tr>
<td align="left">gga</td>
<td align="left">git gui citool --amend</td>
</tr>
<tr>
<td align="left">ggf</td>
<td align="left">git push --force origin $(current_branch)</td>
</tr>
<tr>
<td align="left">ghh</td>
<td align="left">git help</td>
</tr>
<tr>
<td align="left">ggpull</td>
<td align="left">ggl</td>
</tr>
<tr>
<td align="left">ggpur</td>
<td align="left">ggu</td>
</tr>
<tr>
<td align="left">ggpush</td>
<td align="left">ggp</td>
</tr>
<tr>
<td align="left">ggsup</td>
<td align="left">git branch --set-upstream-to = origin/$(current_branch)</td>
</tr>
<tr>
<td align="left">gpsup</td>
<td align="left">git push --set-upstream origin $(current_branch)</td>
</tr>
<tr>
<td align="left">gignore</td>
<td align="left">git update-index --assume-unchanged</td>
</tr>
<tr>
<td align="left">gignored</td>
<td align="left">git ls-files -v | grep "^<a class="internal absent" href="/robbyrussell/oh-my-zsh/wiki/%3Alower%3A">:lower:</a>"</td>
</tr>
<tr>
<td align="left">git-svn-dcommit-push</td>
<td align="left">git svn dcommit &amp;&amp; git push github master:svntrunk</td>
</tr>
<tr>
<td align="left">gk</td>
<td align="left">\gitk --all --branches</td>
</tr>
<tr>
<td align="left">gke</td>
<td align="left">\gitk --all $(git log -g --pretty = format:%h)</td>
</tr>

<tr>
<td align="left">glg</td>
<td align="left">git log --stat --color</td>
</tr>
<tr>
<td align="left">glgg</td>
<td align="left">git log --graph --color</td>
</tr>
<tr>
<td align="left">glgga</td>
<td align="left">git log --graph --decorate --all</td>
</tr>
<tr>
<td align="left">glgm</td>
<td align="left">git log --graph --max-count = 10</td>
</tr>
<tr>
<td align="left">glgp</td>
<td align="left">git log --stat --color -p</td>
</tr>
<tr>
<td align="left">glo</td>
<td align="left">git log --oneline --decorate --color</td>
</tr>
<tr>
<td align="left">glog</td>
<td align="left">git log --oneline --decorate --color --graph</td>
</tr>
<tr>
<td align="left">glol</td>
<td align="left">git log --graph --pretty = format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)&lt;%an&gt;%Creset' --abbrev-commit</td>
</tr>
<tr>
<td align="left">glp</td>
<td align="left">_git_log_prettily</td>
</tr>
<tr>
<td align="left">gmt</td>
<td align="left">git mergetool --no-prompt</td>
</tr>
<tr>
<td align="left">gmtvim</td>
<td align="left">git mergetool --no-prompt --tool = vimdiff</td>
</tr>
<tr>
<td align="left">gmum</td>
<td align="left">git merge upstream/master</td>
</tr>
<tr>
<td align="left">gpd</td>
<td align="left">git push --dry-run</td>
</tr>
<tr>
<td align="left">gpoat</td>
<td align="left">git push origin --all &amp;&amp; git push origin --tags</td>
</tr>
<tr>
<td align="left">gpristine</td>
<td align="left">git reset --hard &amp;&amp; git clean -dfx</td>
</tr>
<tr>
<td align="left">gpu</td>
<td align="left">git push upstream</td>
</tr>
<tr>
<td align="left">gpv</td>
<td align="left">git push -v</td>
</tr>
<tr>
<td align="left">gr</td>
<td align="left">git remote</td>
</tr>
<tr>
<td align="left">gra</td>
<td align="left">git remote add</td>
</tr>
<tr>
<td align="left">grb</td>
<td align="left">git rebase</td>
</tr>
<tr>
<td align="left">grba</td>
<td align="left">git rebase --abort</td>
</tr>
<tr>
<td align="left">grbc</td>
<td align="left">git rebase --continue</td>
</tr>
<tr>
<td align="left">grbi</td>
<td align="left">git rebase -i</td>
</tr>
<tr>
<td align="left">grbm</td>
<td align="left">git rebase master</td>
</tr>
<tr>
<td align="left">grbs</td>
<td align="left">git rebase --skip</td>
</tr>
<tr>
<td align="left">grh</td>
<td align="left">git reset HEAD</td>
</tr>
<tr>
<td align="left">grhh</td>
<td align="left">git reset HEAD --hard</td>
</tr>
<tr>
<td align="left">grmv</td>
<td align="left">git remote rename</td>
</tr>
<tr>
<td align="left">grrm</td>
<td align="left">git remote remove</td>
</tr>
<tr>
<td align="left">grset</td>
<td align="left">git remote set-url</td>
</tr>
<tr>
<td align="left">grt</td>
<td align="left">cd $(git rev-parse --show-toplevel || echo ".")</td>
</tr>
<tr>
<td align="left">gru</td>
<td align="left">git reset --</td>
</tr>
<tr>
<td align="left">grup</td>
<td align="left">git remote update</td>
</tr>
<tr>
<td align="left">grv</td>
<td align="left">git remote -v</td>
</tr>
<tr>
<td align="left">gsb</td>
<td align="left">git status -sb</td>
</tr>
<tr>
<td align="left">gsd</td>
<td align="left">git svn dcommit</td>
</tr>
<tr>
<td align="left">gsi</td>
<td align="left">git submodule init</td>
</tr>
<tr>
<td align="left">gsps</td>
<td align="left">git show --pretty = short --show-signature</td>
</tr>
<tr>
<td align="left">gsr</td>
<td align="left">git svn rebase</td>
</tr>
<tr>
<td align="left">gss</td>
<td align="left">git status -s</td>
</tr>
<tr>
<td align="left">gst</td>
<td align="left">git status</td>
</tr>
<tr>
<td align="left">gsta</td>
<td align="left">git stash save</td>
</tr>
<tr>
<td align="left">gstaa</td>
<td align="left">git stash apply</td>
</tr>
<tr>
<td align="left">gstd</td>
<td align="left">git stash drop</td>
</tr>
<tr>
<td align="left">gstl</td>
<td align="left">git stash list</td>
</tr>
<tr>
<td align="left">gstp</td>
<td align="left">git stash pop</td>
</tr>
<tr>
<td align="left">gstc</td>
<td align="left">git stash clear</td>
</tr>
<tr>
<td align="left">gsts</td>
<td align="left">git stash show --text</td>
</tr>
<tr>
<td align="left">gsu</td>
<td align="left">git submodule update</td>
</tr>
<tr>
<td align="left">gts</td>
<td align="left">git tag -s</td>
</tr>
<tr>
<td align="left">gunignore</td>
<td align="left">git update-index --no-assume-unchanged</td>
</tr>
<tr>
<td align="left">gunwip</td>
<td align="left">git log -n 1 | grep -q -c "--wip--" &amp;&amp; git reset HEAD~1</td>
</tr>
<tr>
<td align="left">gup</td>
<td align="left">git pull --rebase</td>
</tr>
<tr>
<td align="left">gupv</td>
<td align="left">git pull --rebase -v</td>
</tr>
<tr>
<td align="left">gupa</td>
<td align="left">git pull --rebase --autostash</td>
</tr>
<tr>
<td align="left">gupav</td>
<td align="left">git pull --rebase --autostash -v</td>
</tr>
<tr>
<td align="left">gunignore</td>
<td align="left">git update-index --no-assume-unchanged</td>
</tr>
<tr>
<td align="left">glum</td>
<td align="left">git pull upstream master</td>
</tr>
<tr>
<td align="left">gvt</td>
<td align="left">git verify-tag</td>
</tr>
<tr>
<td align="left">gwch</td>
<td align="left">git whatchanged -p --abbrev-commit --pretty = medium</td>
</tr>
<tr>
<td align="left">gwip</td>
<td align="left">git add -A; git rm $(git ls-files --deleted) 2&gt; /dev/null; git commit -m "--wip--"</td>
</tr>



<tr>
<td align="left">gbda</td>
<td align="left">git branch --merged | command grep -vE "^(*|\s<em>master\s</em>$)" | command xargs -n 1 git branch -d</td>
</tr>
<tr>
<td align="left">gbl</td>
<td align="left">git blame -b -w</td>
</tr>
<tr>
<td align="left">gbnm</td>
<td align="left">git branch --no-merged</td>
</tr>
<tr>
<td align="left">gbr</td>
<td align="left">git branch --remote</td>
</tr>

<tr>
<td align="left">gbs</td>
<td align="left">git bisect</td>
</tr>
<tr>
<td align="left">gbsb</td>
<td align="left">git bisect bad</td>
</tr>
<tr>
<td align="left">gbsg</td>
<td align="left">git bisect good</td>
</tr>
<tr>
<td align="left">gbsr</td>
<td align="left">git bisect reset</td>
</tr>
<tr>
<td align="left">gbss</td>
<td align="left">git bisect start</td>
</tr>
<tr>
<td align="left">gca</td>
<td align="left">git commit -v -a</td>
</tr>
<tr>
<td align="left">gcam</td>
<td align="left">git commit -a -m</td>
</tr>
<tr>
<td align="left">gca!</td>
<td align="left">git commit -v -a --amend</td>
</tr>
<tr>
<td align="left">gcan!</td>
<td align="left">git commit -v -a -s --no-edit --amend</td>
</tr>
</tbody>
</table>