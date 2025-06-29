package usecase

import "github.com/brkss/dextrace/internal/domain"






type NightscoutUsecase struct {
    nightscoutRepo domain.NightscoutRepository
}

func NewNightscoutUseCase(nightscoutRepo domain.NightscoutRepository) *NightscoutUsecase {
    return &NightscoutUsecase{
        nightscoutRepo,
    }
}

func (uc *NightscoutUsecase)PushData(data []domain.GetDataResponse) (error) {

    err := uc.nightscoutRepo.PushData(data)
    if err != nil {
        return err;
    }
    return nil
}