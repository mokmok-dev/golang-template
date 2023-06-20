variable "github" {
  type = object({
    owner = string
  })
  default = {
    owner = "mokmok-dev"
  }
}
