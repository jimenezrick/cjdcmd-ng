/*
 * You may redistribute this program and/or modify it under the terms of
 * the GNU General Public License as published by the Free Software Foundation,
 * either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"github.com/inhies/go-cjdns/admin"
	"github.com/spf13/cobra"
	"os"
)

var Pretty bool

func init() {
	DumpCmd.Flags().BoolVarP(&Pretty, "pretty", "p", false, "pretty output")
}

func dumpCmd(cmd *cobra.Command, args []string) {
	c := Connect()
	table, err := c.NodeStore_dumpTable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if Pretty {
		dumpTablePretty(table)
	} else {
		dumpTablePlain(table)
	}
}

func dumpTablePlain(table admin.Routes) {
	table.SortByQuality()
	k := 1
	for _, v := range table {
		if v.Link >= 1 {
			fmt.Fprintf(os.Stdout,
				"%03d IP: %-39v -- Version: %d -- Path: %s -- Link: %s\n",
				k, v.IP, v.Version, v.Path, v.Link)
			k++
		}
	}
}

func dumpTablePretty(table admin.Routes) {
	table.SortByPath()

	fmt.Printf("%s┐\n", table[0].Path)

	printPrettySubtable(table[1:], "")
}

func printPrettySubtable(table admin.Routes, spacer string) {
	var level []admin.Routes

	for i, here := range table {
		if here == nil {
			continue
		}
		// Hit each entry once
		table[i] = nil

		// make a subtable with here at the front
		sublevel := make(admin.Routes, 1)
		sublevel[0] = here
		for j, there := range table {
			if there == nil {
				continue
			}

			if there.Path.IsBehind(*here.Path) {
				sublevel = append(sublevel, there)
				table[j] = nil
			}
		}

		level = append(level, sublevel)
	}

	for _, sublevel := range level[:len(level)-1] {
		here := sublevel[0]
		if len(sublevel) == 1 {
			fmt.Printf("%s%s└─%s\n", here.Path, spacer, here.IP)
			return
		}

		fmt.Printf("%s%s├┬%s\n", here.Path, spacer, here.IP)
		printPrettySubtable(sublevel[1:], spacer+"│")
	}

	sublevel := level[len(level)-1]
	here := sublevel[0]
	if len(sublevel) == 1 {
		fmt.Printf("%s%s└─%s\n", here.Path, spacer, here.IP)
		return
	}
	
	fmt.Printf("%s%s└┬%s\n", here.Path, spacer, here.IP)
	printPrettySubtable(sublevel[1:], spacer+" ")
}
