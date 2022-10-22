package auth

import (
	"testing"
)

func TestUserFromJwt(t *testing.T) {
	user, err := userFromJwt("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoicGRLQmNBYzdsS3pTQzlueWdsQ3dRIn0sImlhdCI6MTY2NjE4NTI2NCwiZXhwIjoxNjY4Nzc3MjY0fQ.BfBhjchBtYOjvXE0_eAydfM5kR4oKPXq2gNUc-XR8mtH3UhR3LZtlGLuFIS2L470hoxmEgH2BfBCMuFnMEzkvuVSKSlOTtCFodlb-Yl49btL_B4fM7eeJZ0n77LFk2aE942AUUAU6_u3y-4dX8-RJobKkeopLQUNRd9zMjSk3w9MrjjgxNqOF2Tj9-1UNuR2McDi2kOBi_1TnaWpM3K8g0ljdeupbMTbEevH2MKbm0IlYZv09e0sWR3rrSDAzaYpc2f8Dqq_ZUP2SBrh_Ly_aIHRJG4gD8788YE-U1GRntpnLXvVh4QnWtTJ5TNFSPlcKGgcBdvK5enLM8MTb2hrag")

	if err != nil{
		t.Error(err)
	} else if user.Id != "pdKBcAc7lKzSC9nyglCwQ"{
		t.Errorf("userFromJwt want test, but: %s", user.Id)
	}

	user2, err := userFromJwt("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoicGRLQmNBYzdsS3pTQzlueWdsQ3dRIn0sImlhdCI6MTY2NjE4NTI2NCwiZXhwIjoxNjY4Nzc3MjY0fQ.BfBhjchBtYOjvXE0_eAydfM5kR4oKPXq2gNUc-XR8mtH3UhR3LZtlGLuFIS2L470hoxmEgH2BfBCMuFnMEzkvuVSKSlOTtCFodlb-Yl49btL_B4fM7eeJZ0n77LFk2aE942AUUAU6_u3y-4dX8-RJobKkeopLQUNRd9zMjSk3w9MrjjgxNqOF2Tj9-1UNuR2McDi2kOBi_1TnaWpM3K8g0ljdeupbMTbEevH2MKbm0IlYZv09e0sWR3rrSDAzaYpc2f8Dqq_ZUP2SBrh_Ly_aIHRJG4gD8788YE-U1GRntpnLXvVh4QnWtTJ5TNFSPlcKGgcBdvK5enLM8MTb2hrag")
	if err != nil{
		t.Error(err)
	} else if user2.Id != "pdKBcAc7lKzSC9nyglCwQ"{
		t.Errorf("userFromJwt want test, but: %s", user2.Id)
	}
}