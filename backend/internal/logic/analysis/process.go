package analysis

// import (
// 	"backend/internal/dao"
// 	"backend/internal/model/do"
// 	"backend/internal/model/entity"
// 	"backend/internal/service"
// 	"context"
// 	"strings"

// 	"github.com/gogf/gf/v2/database/gdb"
// 	"github.com/gogf/gf/v2/util/gconv"
// 	"github.com/google/pprof/profile"
// )

// func (s *sAnalysis) ProcessProfile(ctx context.Context, p *profile.Profile) (err error) {
// 	// 事务处理
// 	return dao.Profile.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
// 		// Step 1: 插入主表 profile
// 		profileId, err := service.Analysis().InsertProfile(ctx, p)
// 		if err != nil {
// 			return err
// 		}
// 		// Step 2: 插入 valueType (SampleType 和 PeriodType)
// 		err = service.Analysis().InsertSampleTypes(ctx, p.SampleType, profileId)
// 		if err != nil {
// 			return err
// 		}
// 		// Step 3: 插入 mapping 信息
// 		mappingIDMap, err := service.Analysis().InsertMappings(ctx, p.Mapping, profileId)
// 		if err != nil {
// 			return err
// 		}

// 		// Step 4: 插入 function 信息
// 		functions, err := service.Analysis().InsertFunctions(ctx, profileId, p.Function)
// 		if err != nil {
// 			return err
// 		}
// 		// Step 5: 插入 location 和 line_entry
// 		locations, err := service.Analysis().InsertLocations(ctx, profileId, p.Location, mappingIDMap)
// 		if err != nil {
// 			return err
// 		}
// 		err = service.Analysis().InsertLineEntries(ctx, p.Location, locations, functions)
// 		if err != nil {
// 			return err
// 		}
// 		// Step 6: 插入 sample 及其关联数据
// 		err = service.Analysis().InsertSamples(ctx, profileId, p.Sample, locations)
// 		if err != nil {
// 			return err
// 		}
// 		// Step 7: 插入注释
// 		err = service.Analysis().InsertComments(ctx, profileId, p.Comments)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }

// // InsertProfile 插入 profile 主表数据
// func (s *sAnalysis) InsertProfile(ctx context.Context, p *profile.Profile) (profileId int64, err error) {
// 	profileId, err = dao.Profile.Ctx(ctx).InsertAndGetId(do.Profile{
// 		DefaultSampleType: p.DefaultSampleType,
// 		DocUrl:            p.DocURL,
// 		DropFrames:        p.DropFrames,
// 		KeepFrames:        p.KeepFrames,
// 		TimeNanos:         p.TimeNanos,
// 		DurationNanos:     p.DurationNanos,
// 		Period:            p.Period,
// 		PeriodTypeType:    p.PeriodType.Type,
// 		PeriodTypeUnit:    p.PeriodType.Unit,
// 	})
// 	if err != nil {
// 		return
// 	}
// 	return profileId, nil
// }

// // InsertSampleTypes 插入 SampleType
// func (s *sAnalysis) InsertSampleTypes(ctx context.Context, types []*profile.ValueType, profileId int64) (err error) {
// 	for _, vt := range types {
// 		if _, err := dao.SampleType.Ctx(ctx).Insert(entity.SampleType{
// 			ProfileId: gconv.Uint64(profileId),
// 			Type:      vt.Type,
// 			Unit:      vt.Unit,
// 		}); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // InsertMappings 插入 mapping 信息
// func (s *sAnalysis) InsertMappings(ctx context.Context, mappings []*profile.Mapping, profileId int64) (mappingIDMap map[uint64]int64, err error) {

// 	mappingIDMap = make(map[uint64]int64)

// 	for _, m := range mappings {
// 		id, err := dao.Mapping.Ctx(ctx).InsertAndGetId(
// 			entity.Mapping{
// 				ProfileId:       gconv.Uint64(profileId),
// 				Start:           m.Start,
// 				Limit:           m.Limit,
// 				Offset:          m.Offset,
// 				File:            m.File,
// 				BuildId:         m.BuildID,
// 				HasFunctions:    gconv.Int(m.HasFunctions),
// 				HasFilenames:    gconv.Int(m.HasFilenames),
// 				HasInlineFrames: gconv.Int(m.HasInlineFrames),
// 				HasLineNumbers:  gconv.Int(m.HasLineNumbers),
// 			},
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		mappingIDMap[m.ID] = id
// 	}
// 	return mappingIDMap, nil
// }

// // insertFunctions 插入 Function 并返回 ID 映射表
// func (s *sAnalysis) InsertFunctions(ctx context.Context, profileID int64, funcs []*profile.Function) (map[uint64]int64, error) {
// 	funcIDMap := make(map[uint64]int64)

// 	for _, f := range funcs {
// 		res, err := dao.Function.Ctx(ctx).Insert(entity.Function{
// 			ProfileId:  gconv.Uint64(profileID),
// 			Name:       f.Name,
// 			SystemName: f.SystemName,
// 			Filename:   f.Filename,
// 			StartLine:  f.StartLine,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		id, _ := res.LastInsertId()
// 		funcIDMap[f.ID] = id
// 	}
// 	return funcIDMap, nil
// }

// // insertLocations 插入 Location 并返回 ID 映射表
// func (s *sAnalysis) InsertLocations(ctx context.Context, profileID int64, locs []*profile.Location, mappingIDMap map[uint64]int64) (map[uint64]int64, error) {
// 	locationIDMap := make(map[uint64]int64)
// 	for _, loc := range locs {
// 		var mappingID interface{}
// 		if loc.Mapping != nil {
// 			mappingID = mappingIDMap[loc.Mapping.ID]
// 		}
// 		res, err := dao.Location.Ctx(ctx).Insert(entity.Location{
// 			ProfileId: gconv.Uint64(profileID),
// 			MappingId: gconv.Uint64(mappingID),
// 			Address:   loc.Address,
// 			IsFolded:  gconv.Int(loc.IsFolded),
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		id, _ := res.LastInsertId()
// 		locationIDMap[loc.ID] = id
// 	}
// 	return locationIDMap, nil
// }

// // insertLineEntries 插入 LineEntry
// func (s *sAnalysis) InsertLineEntries(ctx context.Context, locs []*profile.Location, locationIDMap, funcIDMap map[uint64]int64) error {

// 	for _, loc := range locs {
// 		locID := locationIDMap[loc.ID]
// 		for _, line := range loc.Line {
// 			var funcID interface{}
// 			if line.Function != nil {
// 				funcID = funcIDMap[line.Function.ID]
// 			}
// 			if _, err := dao.LineEntry.Ctx(ctx).Insert(
// 				entity.LineEntry{
// 					LocationId: gconv.Uint64(locID),
// 					Line:       line.Line,
// 					FunctionId: gconv.Uint64(funcID),
// 					Column:     line.Column,
// 				},
// 			); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// // insertSamples 插入 Sample 及关联数据
// func (s *sAnalysis) InsertSamples(ctx context.Context, profileID int64, samples []*profile.Sample, locationIDMap map[uint64]int64) error {

// 	for _, sample := range samples {
// 		// 插入 sample 表
// 		sample_Value := strings.Join(gconv.Strings(sample.Value), ",")
// 		sampleID, err := dao.Sample.Ctx(ctx).InsertAndGetId(entity.Sample{
// 			ProfileId: gconv.Uint64(profileID),
// 			Value:     sample_Value,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// 插入 sample_Location（反向顺序，从根到叶子）
// 		for orderIdx, loc := range sample.Location {
// 			locID := locationIDMap[loc.ID]
// 			if _, err := dao.SampleLocation.Ctx(ctx).Insert(
// 				ctx,
// 				entity.SampleLocation{
// 					SampleId:   gconv.Uint64(sampleID),
// 					LocationId: gconv.Uint64(locID),
// 					OrderIndex: gconv.Uint(orderIdx),
// 				},
// 			); err != nil {
// 				return err
// 			}
// 		}

// 		// 插入 sample_Label
// 		for key, values := range sample.Label {
// 			for _, val := range values {
// 				if _, err := dao.SampleLabel.Ctx(ctx).Insert(
// 					entity.SampleLabel{
// 						SampleId: gconv.Uint64(sampleID),
// 						KeyStr:   key,
// 						ValueStr: val,
// 					},
// 				); err != nil {
// 					return err
// 				}
// 			}
// 		}

// 		// 插入 sample_numLabel
// 		for key, values := range sample.NumLabel {
// 			units := sample.NumUnit[key]
// 			for i, val := range values {
// 				unit := ""
// 				if i < len(units) {
// 					unit = units[i]
// 				}
// 				if _, err := dao.SampleNumlabel.Ctx(ctx).Insert(
// 					entity.SampleNumlabel{
// 						SampleId:   gconv.Uint64(sampleID),
// 						KeyStr:     key,
// 						ValueIndex: gconv.Uint(i),
// 						NumValue:   val,
// 						UnitStr:    unit,
// 					},
// 				); err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// // insertComments 插入注释
// func (s *sAnalysis) InsertComments(ctx context.Context, profileID int64, comments []string) error {

// 	for i, comment := range comments {
// 		if _, err := dao.ProfileComment.Ctx(ctx).Insert(
// 			entity.ProfileComment{
// 				ProfileId:  gconv.Uint64(profileID),
// 				Comment:    comment,
// 				OrderIndex: gconv.Uint(i),
// 			},
// 		); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
