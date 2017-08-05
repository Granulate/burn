// Copyright © 2017 Martin Spier <spiermar@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stacko

func ParseFolded(filename string) Profile {
	rootNode := Node{"root", 0, make(map[string]*Node)}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		sep := strings.LastIndex(line, " ")

		s := line[:sep]
		v := line[sep+1:]

		stack := strings.Split(s, ";")
		sort.Sort(sort.Reverse(sort.StringSlice(stack)))

		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		rootNode.Add(&stack, len(stack)-1, i)
	}

	if err := scanner.Err(); err != nil {

	}

	profile := Profile{rootNode, []string{}, ""}

	return profile
}