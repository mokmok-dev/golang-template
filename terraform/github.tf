resource "github_repository" "golang_template" {
  name        = "golang-template"
  description = "Server-side application template written in Go"

  allow_auto_merge       = true
  allow_merge_commit     = false
  allow_rebase_merge     = false
  allow_squash_merge     = true
  allow_update_branch    = true
  delete_branch_on_merge = true
  has_issues             = true
  is_template            = true
  visibility             = "public"
}

resource "github_branch_protection" "main" {
  repository_id       = github_repository.golang_template.id
  pattern             = "main"
  allows_force_pushes = true
}
