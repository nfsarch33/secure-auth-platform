# How to automate tagging and release workflows in GitHub


**Note**

This guide explains this concept in vanilla Git. For Graphite documentation, see our [CLI docs](https://graphite.com/docs/command-reference?utm_source=guidesCallout).

Automating tagging and release workflows in GitHub can significantly streamline the process of deploying software, ensuring consistency and reliability while reducing the likelihood of human error. This guide explores how to set up automation for tagging and managing releases in GitHub, using built-in features like GitHub Actions.

### Benefits of automating tagging and release workflows

Automating these workflows offers several advantages:

* **Consistency** : Automation ensures that every release follows the same steps and standards, reducing the chances of mistakes.
* **Efficiency** : Reduces the time and effort required to prepare and execute releases.
* **Traceability** : Automatic tagging and releasing help in maintaining a clear record of changes, facilitating easier rollback and history tracking.

### Setting up automated tagging and release workflows in GitHub

#### Step 1: Define the workflow

Before automating the process, clearly define the conditions under which a new tag is created and a release is made. Common triggers include:

* Merging a pull request into the `main` branch.
* Pushing a commit with a specific commit message format.
* Manually triggering via GitHub's UI when a release candidate is ready.

#### Step 2: Configure GitHub Actions for tagging

GitHub Actions can automate the creation of tags based on your triggers. Here's how you can set up an action to tag commits:

1. **Create a new GitHub Actions workflow file** in your repository under `.github/workflows`, for example, `tagging.yml`.
2. **Define the workflow trigger** . For example, trigger on push (a merge will also trigger a push event) to the `main` branch:

Terminal

<pre class="prism-code language-bash console_pre__3Tuei"><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">name: Automated Tagging</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">on:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">  push:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    branches:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">      - main</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">
</span></span></div></pre>

3. **Add steps to create a tag** . You can use the [Git CLI](https://graphite.com/guides/git-push-tag) or other actions available in the [marketplace](https://github.com/marketplace?query=tag):

Terminal

<pre class="prism-code language-bash console_pre__3Tuei"><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">jobs:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">  tag:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    runs-on: ubuntu-latest</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    steps:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">      - uses: actions/checkout@v2</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">      - name: Create Tag</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">        run: </span><span class="token operator">|</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain"></span><span class="token function">git</span><span class="token plain"> config --local user.email </span><span class="token string">"action@github.com"</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain"></span><span class="token function">git</span><span class="token plain"> config --local user.name </span><span class="token string">"GitHub Action"</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain"></span><span class="token function">git</span><span class="token plain"> tag -a </span><span class="token string">"v</span><span class="token string variable">${{ github.run_number }</span><span class="token string">}"</span><span class="token plain"> -m </span><span class="token string">"Release v</span><span class="token string variable">${{ github.run_number }</span><span class="token string">}"</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain"></span><span class="token function">git</span><span class="token plain"> push origin </span><span class="token string">"v</span><span class="token string variable">${{ github.run_number }</span><span class="token string">}"</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">
</span></span></div></pre>

This script configures git, creates a tag based on the run number of the workflow, and pushes it back to the repository.

#### Step 3: Automate GitHub releases

Once a tag is created, you can then automate the release process:

1. **Extend the GitHub Actions workflow** to include release creation after tagging:

Terminal

<pre class="prism-code language-bash console_pre__3Tuei"><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">- name: Create Release</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">  uses: actions/create-release@v1</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">  env:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    GITHUB_TOKEN: </span><span class="token variable">${{ secrets.GITHUB_TOKEN }</span><span class="token punctuation">}</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">  with:</span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    tag_name: </span><span class="token string">'v${{ github.run_number }}'</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    release_name: </span><span class="token string">'Release v${{ github.run_number }}'</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">    body: </span><span class="token string">'Description of the changes in this release'</span><span class="token plain"></span></span></div><div class="token-line console_line__3dU_m"><span class="console_line-content__YE6LA"><span class="token plain">
</span></span></div></pre>

As an example, this GitHub Actions step creates a new release in the repository, automatically names and tags it based on the workflow's run number, and includes a description of the changes in the release body.

2. **Customize the release step** by specifying the tag name, release name, and body of the release note.

By automating the tagging and release workflows in GitHub, you not only save time and effort but also improve the overall reliability of your deployment process. This setup ensures that all releases are consistent and traceable, providing a robust framework for managing software deployments efficiently and effectively.
