export def get_attributes [
    resource_type: string
    resource_name: string
    state_file: string = "terraform.tfstate"
] {
    open $state_file
    | from json
    | get resources
    | where type == $resource_type and name == $resource_name
    | get instances
    | flatten
    | get attributes
    | first
}
