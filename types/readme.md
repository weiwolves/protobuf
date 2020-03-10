# Initialize protobuf types

To use the protobuf types, they have to be linked to the ```/usr/local/include```
folder, in the same way like Google's protobuf well-known-types.

```console
foo@bar:~$ mkdir /usr/local/include/weiwolves
foo@bar:~$ ln -s $GOPATH/src/github.com/weiwolves/protobuf/types /usr/local/include/weiwolves/protobuf
```

Now the protobuf types can be imported into your proto definition file.
```
syntax = "proto3";

package startrek;

import "weiwolves/protobuf/timestamp.proto";
import "weiwolves/protobuf/nullstring.proto";
import "weiwolves/protobuf/nullint.proto";

message StarfleetShip {
	string name = 1;
	weiwolves.protobuf.NullInt passengers = 2;
	weiwolves.protobuf.NullString mission = 3;
	weiwolves.protobuf.Timestamp departure_time = 4;
}
```
