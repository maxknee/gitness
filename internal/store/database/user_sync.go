// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package database

import (
	"context"

	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/internal/store/database/mutex"
	"github.com/harness/gitness/types"
)

var _ store.UserStore = (*UserStoreSync)(nil)

// NewUserStoreSync returns a new UserStoreSync.
func NewUserStoreSync(store *UserStore) *UserStoreSync {
	return &UserStoreSync{base: store}
}

// UserStoreSync synronizes read and write access to the
// user store. This prevents race conditions when the database
// type is sqlite3.
type UserStoreSync struct{ base *UserStore }

// Find finds the user by id.
func (s *UserStoreSync) Find(ctx context.Context, id int64) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.base.Find(ctx, id)
}

// FindUID finds the user by uid.
func (s *UserStoreSync) FindUID(ctx context.Context, uid string) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.base.FindUID(ctx, uid)
}

// FindEmail finds the user by email.
func (s *UserStoreSync) FindEmail(ctx context.Context, email string) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.base.FindEmail(ctx, email)
}

// List returns a list of users.
func (s *UserStoreSync) List(ctx context.Context, opts *types.UserFilter) ([]*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.base.List(ctx, opts)
}

// Create saves the user details.
func (s *UserStoreSync) Create(ctx context.Context, user *types.User) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Create(ctx, user)
}

// Update updates the user details.
func (s *UserStoreSync) Update(ctx context.Context, user *types.User) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Update(ctx, user)
}

// Delete deletes the user.
func (s *UserStoreSync) Delete(ctx context.Context, id int64) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Delete(ctx, id)
}

// Count returns a count of users.
func (s *UserStoreSync) Count(ctx context.Context) (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()
	return s.base.Count(ctx)
}