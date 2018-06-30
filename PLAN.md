# ABTesting Software

Software to help developers test features, allowing them to turn on and off 
certain features depending on feature flags provided by a language specific 
SDK connecting to a server.  Spec comes below.

## Models

#### User 
- ID: Unique integer
- KEY: unique identifier string
- NAME: Human readable string
- ATTRIBUTES: Key-Value Stores, set by program

Users objects holds objects that are associated with both groups and 
experiments.  If associated with a group, then it inherits all the functionality
of a group (permissions of a feature essentially).  If it is part of an 
experiment, it will be chosen by the experiment's random percentage value to 
determine if it receives a feature.

#### Group
- ID: Unique integer
- KEY: unique identifier string
- NAME: Human readable string
- ATTRIBUTES: Key-Value Stores, set by program.

Groups users that are affected by the same experiments.  Sample usage would 
include having a bunch of beta users, or a development group only.  If a group 
is associated with an Experiment, then it will automatically get those features.


#### Experiment: 
- ID: unique integer
- KEY: unique identifier string
- NAME: Human readable string
- PERCENTAGE: Specify which % of users to roll out a feature to.  Float.
- PROJECT: Which project the experiment is part of.  Foreign Key.
- ATTRIBUTES: Key-Value Stores, set by program.

Denotes a feature that you want to test, whether it be a static page or the 
color of a button. Can set percentage of users that see Features A, B, ..., 
etc.  Each Experiment is associated with one Project.

#### Project 
- ID: unique integer
- KEY: unique identifier string
- NAME: Human readable string
- ATTRIBUTES: Key-Value Stores, set by program.

Project houses groups of experiments.  Typical usage is each software
project is it's own Project, and each Project contains lots of experiments.

#### Relationships of models
- Users to Groups: Many to Many
    - Users can belong to many groups
    - Groups can have any number of people
- User to Experiments: Many to Many
    - Users can belong to multiple Experiments
    - Experiments can have multiple Users
- Groups to Experiments: Many to Many
    - Groups can belong to multiple Experiments
    - Experiments can have multiple Groups
- Experiments to Projects: Many to One
    - Experiments can belong to one Project
    - Project can have many Experiments

Thus many to many relationships will require a joining table.

## Services

#### Project
- Get All Projects
    - /projects GET
- Get Project
    - /projects?project=name GET 
- Create Project
    - /projects POST
- Update Project
    - /projects?project=name PUT
- Delete Project
    - /projects?project=name DELETE
- Get All Experiments
    - /projects/experiments?project=name GET

#### Experiment
- Get All Experiments
    - /experiments?project=name GET
- Get Experiment
    - /experiments?project=name&experiment=name GET
- Create Experiment
    - /experiments?project=name POST
- Update Experiment
    - /experiments?project=name&experiment=name PUT
- Delete Experiment
    - /experiments?project=name&experiment=name DELETE
- Get all groups associated
    - /experiments/groups?project=name&experiment=name GET
- Get all users associated
    - /experiments/users?project=name&experiment=name GET
- Add a group to an experiment
    - /experiments/groups?project=name&experiment=name POST
- Add a user to an experiment
    - /experiments/users?project=name&experiment=name POST
- Remove a group from an experiment
    - /experiments/groups?project=name&experiment=name&group=name DELETE
- Remove a user from an experiment
    - /experiments/users?project=name&experiment=name&user=name POST

#### Groups
- Get All Groups
    - /groups GET
- Get Groups
    - /groups?group=name GET
- Create Group
    - /groups POST
- Update Group
    - /groups?group=name PUT
- Delete Group
    - /groups?group=name DELETE
- Get all Users associated
    - /groups/users?group=name GET
- Get all Experiments group is in
    - /groups/experiments?group=name GET


#### Users
- GET All Users
    - /users GET
- Get User
    - /users?user=name
- Create User
    - /users POST
- Update User
    - /users?user=name PUT
- Delete User
    - /users?user=name DELETE
- Get all Groups associated
    - /users/groups?user=name GET
- Get all Experiments associated
    - /users/experiments?user=name GET

## External Services and Dependencies
#### Dep
Dep is dependency management, for Go packages.  Considering switching to vgo, if proposal is accepted and widely used.

#### Echo
Echo server/router, provides valuable middleware to save time.

#### Postgres
Postgres for database.  Most open source, fully featured ACID compliant database.

#### Docker
Should provide a containerized version of the database.

#### Logging
Could use Echo logger, or some other logger.  Undecided. Options below
- go.uber.org/zap
- github.com/golang/glog
- github.com/sirupsen/logrus
