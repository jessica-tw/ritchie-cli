/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package modifier

import (
	"net/http"
	"testing"

	"github.com/ZupIT/ritchie-cli/pkg/git/github"
	"github.com/stretchr/testify/assert"
)

func TestTemplateRelease(t *testing.T) {
	repoInfo := github.NewRepoInfo(TemplateFormulasRepoURL, "")
	githubRepo := github.NewRepoManager(http.DefaultClient)
	tag, _ := githubRepo.LatestTag(repoInfo)

	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "modify with success",
			args: args{
				b: []byte(`{tag}`),
			},
			want: []byte(tag.Name),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TemplateRelease{}
			got := tr.modify(tt.args.b)
			assert.Equal(t, got, tt.want)
		})
	}
}
