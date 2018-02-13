/*
 *     Copyright (C) 2018  Ontario Institute for Cancer Research
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package endpoint 

import (
	"net/url"
	"strings"
	"path"
)

type Endpoint struct {
	BaseURL *url.URL
}

func (s *Endpoint) makeURL(args ...string) url.URL {
	requestURL := *s.BaseURL
   	requestURL.Path = path.Join(s.BaseURL.Path, path.Join(args...)) 
        return requestURL
}

// Upload uploads the file contents and returns the response
func (s *Endpoint) Upload(studyID string, async bool) url.URL {
        var url = s.makeURL("upload", studyID)
        if (async) {
	  url.Path = path.Join(url.Path, "async")
	}
	return url
}

// GetStatus return the status JSON of an uploadID
func (s *Endpoint) GetStatus(studyID string, uploadID string) url.URL {
        return s.makeURL("upload",studyID,"status",uploadID)
}

func (s *Endpoint) IsAlive() url.URL {
	return s.makeURL("isAlive")
}

func truth(condition bool) string {
	if condition {
		return "true"
	} 
	return "false"
}

// Save saves the specified uploadID as an analysis assuming it had passed validation
func (s *Endpoint) Save(studyID string, uploadID string, ignoreCollisions bool) url.URL {
	
	return s.makeURL("upload", studyID, "save", uploadID,"?ignoreAnalysisIDCollisions=" + truth(ignoreCollisions) )
}

// Publish publishes a specified saved analysisID
func (s *Endpoint) Publish(studyID string, analysisID string) url.URL {
	return s.makeURL("studies", studyID, "publish", analysisID)
}

func (s *Endpoint) Suppress(studyID string, analysisID string) url.URL {
	return s.makeURL("studies", studyID,"analysis", "suppress", analysisID)
}

func (s *Endpoint) GetAnalysis(studyID string, analysisID string) url.URL {
	return s.makeURL("studies",studyID, "analysis",analysisID)
}

func (s *Endpoint) GetAnalysisFiles(studyID string, analysisID string) url.URL {
	return s.makeURL("studies",studyID, "analysis",analysisID,"files")
}

func (s *Endpoint) IdSearch(studyID string, searchParams string) url.URL {
	return s.makeURL("studies", studyID, "analysis/search/id?", 
		searchParams)
}

func (s *Endpoint) InfoSearch(studyID string, includeInfo bool, searchTerms []string) url.URL {
	var params = "info?includeInfo=" + truth(includeInfo) + "&" + strings.Join(searchTerms,"&")
	return s.makeURL("studies", studyID, "search", params) 
} 
