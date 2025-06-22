package knowledgerepo

import(
	"fmt"
	"strings"

	"unimatch/data/knowledge"
	"unimatch/data/network"
)


type KnowledgeRepository interface {
	LoadKnowledgeBase() (int, error)
}

type knowledgeRepositoryImpl struct {
	knowledgeSource knowledgesource.Knowledge
	api api.API
}

var homeRepository KnowledgeRepository

func GetKnowledgeRepo() KnowledgeRepository {
	if homeRepository == nil {
		homeRepository = knowledgeRepositoryImpl{ knowledgeSource: knowledgesource.GetKnowledgeSource(), api: api.GetApi() }
	}

	return homeRepository
}

func (r knowledgeRepositoryImpl) LoadKnowledgeBase() (int, error) {
	matchingCounterPredicates := `member(X, [X|_]).
	member(X, [_|T]) :- member(X, T).

	% --- count how many matches from List1 exist in List2 ---
	count_matches([], _, 0).
	count_matches([H|T], List, Count) :-
	member(H, List) ->
	(count_matches(T, List, Rest), Count is Rest + 1);
	count_matches(T, List, Count).
	`
	r.knowledgeSource.LoadRule(matchingCounterPredicates)

	greaterThanPredicate := "greater_than(X, Y) :- X is Y + 1."

	r.knowledgeSource.LoadRule(greaterThanPredicate)

	careerMatchingRules := `match_all_careers(UserAptitudes, UserSkills, UserInterests, Faculty, Career, AptitudeMatches, SkillMatches, InterestMatches, AptitudeTotal, SkillTotal, InterestTotal) :-
	career_details(Faculty, Career, AList, SList, IList, AptitudeTotal, SkillTotal, InterestTotal),
    count_matches(UserAptitudes, AList, AptitudeMatches),
    count_matches(UserSkills, SList, SkillMatches),
    count_matches(UserInterests, IList, InterestMatches),
	(greater_than(AptitudeMatches, 0) ; 
	greater_than(SkillMatches, 0) ; 
	greater_than(InterestMatches, 0)).`

	r.knowledgeSource.LoadRule(careerMatchingRules)

	apiFacts, err := r.api.GetFacts()

	if err != nil {
		return -1, err
	}

	// fmt.Println("apiFacts")
	// fmt.Println(apiFacts)
	for _, fac := range apiFacts {
		
		factString := fmt.Sprintf(
			"career_details('%s', '%s', %s, %s, %s, %d, %d, %d).",
			fac.Faculty, 
			fac.Career,
			r.formatPrologList(fac.Aptitude),
			r.formatPrologList(fac.Skill), 
			r.formatPrologList(fac.Interest),
			len(fac.Aptitude),
			len(fac.Skill),
			len(fac.Interest),
		)

		r.knowledgeSource.LoadCareerFact(factString)
	}

	// fmt.Println("careerFacts")
	// fmt.Println(apiFacts)

	return len(apiFacts), nil
}

func(r knowledgeRepositoryImpl) formatPrologList(items []string) string {
	var quoted []string

	for _, item := range items {
		str := fmt.Sprintf("%v", item) // safely convert anything to string
		quoted = append(quoted, fmt.Sprintf("'%s'", str))
	}

	return "[" + strings.Join(quoted, ", ") + "]"
}

// func main() {
// 	n, err := GetKnowledgeRepo().LoadKnowledgeBase()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(n)
// }