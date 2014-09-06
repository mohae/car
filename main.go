// Copyright © 2014, All rights reserved
// Joel Scoble, https://github.com/mohae/clitpl
//
// This is licensed under The MIT License. Please refer to the included
// LICENSE file for more information. If the LICENSE file has not been
// included, please refer to the url above.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License
//
// clitpl is a basic implementation of michellh's cli package. clitpl uses
// cli's example, and the implementations used in both mitchellh's and
// hashicorp's implemenations.
//
package main

import (
	"os"
	"runtime"
)

// This is modeled on mitchellh's realmain wrapper
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	os.Exit(appMain())
}
