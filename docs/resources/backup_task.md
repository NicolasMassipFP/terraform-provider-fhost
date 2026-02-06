---
page_title: "smc_backup_task"
subcategory: ""
description: |-
  This represents a Backup Task, which is used to create backups of the system. It is a type of task that can be scheduled and executed to ensure that system data is backed up properly.
---

# smc_backup_task

This represents a Backup Task, which is used to create backups of the system. It is a type of task that can be scheduled and executed to ensure that system data is backed up properly.


## Simple Attributes
    
- `backup_comment` (String) Comment for the backup task, describing its purpose or details.
- `comment` (String) An optional comment for the element. This field is not required.
- `log_data_must_be_saved` (Boolean) Flag to indicate whether log data must be saved during the backup process. This is used to ensure that log data is preserved in the backup.
- `name` (String) Name of the object.
- `password` (String) Password for the backup task, used to secure the backup data. This should be set using the clear password method.
- `resources` (List of String) URI of the resource.
- `scriptToExecute` (String) The script to execute after the backup is completed. This can be used to perform additional actions or cleanup after the backup process.
- `serverPath` (String) The server path where the backup will be stored. This is the destination path for the backup files on the server.


## Readonly Attributes
    
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.