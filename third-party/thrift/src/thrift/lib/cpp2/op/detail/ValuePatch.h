/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
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

#pragma once

#include <utility>

#include <thrift/lib/cpp2/op/detail/BasePatch.h>

namespace apache {
namespace thrift {
namespace op {
namespace detail {

/// A patch adapter that only supports 'assign',
/// which is the minimum any patch should support.
///
/// The `Patch` template parameter must be a Thrift struct with the following
/// fields:
/// * `optional T assign`
template <typename Patch>
class AssignPatch : public BaseAssignPatch<Patch, AssignPatch<Patch>> {
  using Base = BaseAssignPatch<Patch, AssignPatch>;
  using T = typename Base::value_type;

 public:
  using Base::apply;
  using Base::Base;
  using Base::operator=;

  /// @brief This API uses the Visitor pattern to describe how Patch is applied.
  /// For each operation that will be performed by the patch, the corresponding
  /// method (that matches the write API) will be invoked.
  ///
  /// Users should provide a visitor with the following methods
  ///
  ///     struct Visitor {
  ///       void assign(const MyClass&);
  ///     }
  ///
  /// For example:
  ///
  ///     MyClassPatch patch;
  ///     patch = myClass;
  ///
  /// `patch.customVisit(v)` will invoke the following methods
  ///
  ///     v.assign(myClass);
  template <typename Visitor>
  void customVisit(Visitor&& v) const {
    if (auto p = data_.assign()) {
      std::forward<Visitor>(v).assign(*p);
    }
  }

  void apply(T& val) const {
    struct Visitor {
      T& v;
      void assign(const T& t) { v = t; }
    };
    return customVisit(Visitor{val});
  }

 private:
  using Base::data_;
};

/// Patch for a Thrift bool.
///
/// The `Patch` template parameter must be a Thrift struct with the following
/// fields:
/// * `optional T assign`
/// * `terse bool clear`
/// * `terse bool invert`
template <typename Patch>
class BoolPatch : public BaseClearPatch<Patch, BoolPatch<Patch>> {
  using Base = BaseClearPatch<Patch, BoolPatch>;
  using T = typename Base::value_type;

 public:
  using Base::apply;
  using Base::Base;
  using Base::operator=;

  /// Creates a new patch that invert the bool.
  static BoolPatch createInvert() { return !BoolPatch{}; }

  /// Inverts the bool.
  void invert() {
    auto& val = assignOr(*data_.invert());
    val = !val;
  }

  /// @copybrief AssignPatch::customVisit
  ///
  /// Users should provide a visitor with the following methods
  ///
  ///     struct Visitor {
  ///       void assign(bool);
  ///       void clear();
  ///       void invert();
  ///     }
  ///
  /// For example:
  ///
  ///     auto patch = BoolPatch::createClear();
  ///     patch = !patch;
  ///
  /// `patch.customVisit(v)` will invoke the following methods
  ///
  ///     v.clear();
  ///     v.invert();
  template <typename Visitor>
  void customVisit(Visitor&& v) const {
    if (false) {
      // Test whether the required methods exist in Visitor
      v.assign(T{});
      v.clear();
      v.invert();
    }
    if (!Base::template customVisitAssignAndClear(v) && *data_.invert()) {
      std::forward<Visitor>(v).invert();
    }
  }

  void apply(T& val) const {
    struct Visitor {
      T& v;
      void assign(T b) { v = b; }
      void clear() { v = false; }
      void invert() { v = !v; }
    };

    return customVisit(Visitor{val});
  }

 private:
  using Base::assignOr;
  using Base::data_;

  friend BoolPatch operator!(BoolPatch val) { return (val.invert(), val); }
};

/// Patch for a numeric Thrift types.
///
/// The `Patch` template parameter must be a Thrift struct with the following
/// fields:
/// * `optional T assign`
/// * `terse bool clear`
/// * `terse T add`
template <typename Patch>
class NumberPatch : public BaseClearPatch<Patch, NumberPatch<Patch>> {
  using Base = BaseClearPatch<Patch, NumberPatch>;
  using T = typename Base::value_type;
  using Tag = type::infer_tag<T>;

 public:
  using Base::apply;
  using Base::Base;
  using Base::operator=;

  /// Creates a new patch that increases the value.
  template <typename U>
  static NumberPatch createAdd(U&& val) {
    NumberPatch patch;
    patch.add(std::forward<U>(val));
    return patch;
  }
  /// Creates a new patch that decreases the value.
  template <typename U>
  static NumberPatch createSubtract(U&& val) {
    NumberPatch patch;
    patch.subtract(std::forward<U>(val));
    return patch;
  }

  /// Increases the value.
  template <typename U>
  void add(U&& val) {
    assignOr(*data_.add()) += std::forward<U>(val);
  }

  /// Decreases the value.
  template <typename U>
  void subtract(U&& val) {
    assignOr(*data_.add()) -= std::forward<U>(val);
  }

  /// @copybrief AssignPatch::customVisit
  ///
  /// Users should provide a visitor with the following methods
  ///
  ///     struct Visitor {
  ///       void assign(Number);
  ///       void clear();
  ///       void add(Number);
  ///     }
  ///
  /// For example:
  ///
  ///     auto patch = I32Patch::createClear();
  ///     patch += 10;
  ///
  /// `patch.customVisit(v)` will invoke the following methods
  ///
  ///     v.clear();
  ///     v.add(10);
  template <typename Visitor>
  void customVisit(Visitor&& v) const {
    if (false) {
      // Test whether the required methods exist in Visitor
      v.assign(T{});
      v.clear();
      v.add(T{});
    }
    if (!Base::template customVisitAssignAndClear(v)) {
      v.add(*data_.add());
    }
  }

  void apply(T& val) const {
    struct Visitor {
      T& v;
      void assign(const T& t) { v = t; }
      void clear() { v = 0; }
      void add(const T& t) { v += t; }
    };

    return customVisit(Visitor{val});
  }

  /// Increases the value.
  template <typename U>
  NumberPatch& operator+=(U&& val) {
    add(std::forward<U>(val));
    return *this;
  }

  /// Decreases the value.
  template <typename U>
  NumberPatch& operator-=(U&& val) {
    subtract(std::forward<T>(val));
    return *this;
  }

 private:
  using Base::assignOr;
  using Base::data_;

  template <typename U>
  friend NumberPatch operator+(NumberPatch lhs, U&& rhs) {
    lhs.add(std::forward<U>(rhs));
    return lhs;
  }

  template <typename U>
  friend NumberPatch operator+(U&& lhs, NumberPatch rhs) {
    rhs.add(std::forward<U>(lhs));
    return rhs;
  }

  template <typename U>
  friend NumberPatch operator-(NumberPatch lhs, U&& rhs) {
    lhs.subtract(std::forward<U>(rhs));
    return lhs;
  }
};

/// Base class for string/binary patch types.
///
/// The `Patch` template parameter must be a Thrift struct with the following
/// fields:
/// * `optional T assign`
/// * `terse bool clear`
/// * `terse T append`
/// * `terse T prepend`
template <typename Patch, typename Derived>
class BaseStringPatch : public BaseContainerPatch<Patch, Derived> {
  using Base = BaseContainerPatch<Patch, Derived>;
  using T = typename Base::value_type;

 public:
  using Base::Base;
  using Base::operator=;

  /// Creates a patch that prepends a string.
  template <typename U>
  static Derived createPrepend(U&& val) {
    Derived patch;
    patch.prepend(std::forward<U>(val));
    return patch;
  }

  /// Creates a patch that appends a string.
  template <typename... Args>
  static Derived createAppend(Args&&... args) {
    Derived patch;
    patch.append(std::forward<Args>(args)...);
    return patch;
  }

  /// Appends a string.
  template <typename U>
  Derived& operator+=(U&& val) {
    derived().append(std::forward<U>(val));
    return derived();
  }

  /// @copybrief AssignPatch::customVisit
  ///
  /// Users should provide a visitor with the following methods
  ///
  ///     struct Visitor {
  ///       void assign(const String&);
  ///       void clear();
  ///       void prepend(const String&);
  ///       void append(const String&);
  ///     }
  ///
  /// For example:
  ///
  ///     auto patch = StringPatch::createPrepend("(");
  ///     patch += ")";
  ///
  /// `patch.customVisit(v)` will invoke the following methods
  ///
  ///     v.prepend("(");
  ///     v.append(")");
  template <class Visitor>
  void customVisit(Visitor&& v) const {
    if (false) {
      // Test whether the required methods exist in Visitor
      v.assign(T{});
      v.clear();
      v.prepend(T{});
      v.append(T{});
    }
    if (!Base::template customVisitAssignAndClear(std::forward<Visitor>(v))) {
      std::forward<Visitor>(v).prepend(*data_.prepend());
      std::forward<Visitor>(v).append(*data_.append());
    }
  }

 protected:
  using Base::assignOr;
  using Base::data_;
  using Base::derived;

 private:
  /// Concat two strings.
  template <typename U>
  friend Derived operator+(Derived lhs, U&& rhs) {
    return lhs += std::forward<U>(rhs);
  }
  /// Concat two strings.
  template <typename U>
  friend Derived operator+(U&& lhs, Derived rhs) {
    rhs.prepend(std::forward<U>(lhs));
    return rhs;
  }
};

/// Patch for a Thrift string.
///
/// The `Patch` template parameter must be a Thrift struct with the following
/// fields:
/// * `optional string assign`
/// * `terse bool clear`
/// * `terse string append`
/// * `terse string prepend`
template <typename Patch>
class StringPatch : public BaseStringPatch<Patch, StringPatch<Patch>> {
  using Base = BaseStringPatch<Patch, StringPatch>;
  using T = typename Base::value_type;

 public:
  using Base::apply;
  using Base::Base;
  using Base::operator=;

  /// Appends a string.
  template <typename... Args>
  void append(Args&&... args) {
    assignOr(*data_.append()).append(std::forward<Args>(args)...);
  }

  /// Prepends a string.
  template <typename U>
  void prepend(U&& val) {
    T& cur = assignOr(*data_.prepend());
    cur = std::forward<U>(val) + std::move(cur);
  }

  void apply(T& val) const {
    struct Visitor {
      T& v;
      void assign(const T& t) { v = t; }
      void clear() { v = ""; }
      // TODO: Optimize this
      void prepend(const T& t) { v = t + v; }
      void append(const T& t) { v += t; }
    };

    return Base::customVisit(Visitor{val});
  }

 private:
  using Base::assignOr;
  using Base::data_;
};

inline const folly::IOBuf& emptyIOBuf() {
  static const auto empty =
      folly::IOBuf::wrapBufferAsValue(folly::StringPiece(""));
  return empty;
}

/// Patch for a Thrift Binary.
///
/// The `Patch` template parameter must be a Thrift struct with the following
/// fields:
/// * `optional standard.ByteBuffer assign`
/// * `terse bool clear`
/// * `terse standard.ByteBuffer append`
/// * `terse standard.ByteBuffer prepend`
template <typename Patch>
class BinaryPatch : public BaseStringPatch<Patch, BinaryPatch<Patch>> {
  using Base = BaseStringPatch<Patch, BinaryPatch>;
  using T = typename Base::value_type;

 public:
  using Base::apply;
  using Base::Base;
  using Base::operator=;

  /// Appends a binary string.
  template <typename... Args>
  void append(Args&&... args) {
    folly::IOBufQueue queue{folly::IOBufQueue::cacheChainLength()};
    auto& cur = assignOr(*data_.append());
    queue.append(cur);
    queue.append(std::forward<Args>(args)...);
    cur = queue.moveAsValue();
  }

  /// Prepends a binary string.
  template <typename U>
  void prepend(U&& val) {
    auto& cur = assignOr(*data_.prepend());
    folly::IOBufQueue queue{folly::IOBufQueue::cacheChainLength()};
    queue.append(std::forward<U>(val));
    queue.append(cur);
    cur = queue.moveAsValue();
  }

  void apply(T& val) const {
    struct Visitor {
      T& v;
      std::deque<const T*> bufs = {&v};
      void assign(const T& t) { bufs = {&t}; }
      void clear() { bufs = {&emptyIOBuf()}; }
      void prepend(const T& t) { bufs.push_front(&t); }
      void append(const T& t) { bufs.push_back(&t); }
      ~Visitor() {
        folly::IOBufQueue queue{folly::IOBufQueue::cacheChainLength()};
        for (auto buf : bufs) {
          queue.append(*buf);
        }
        v = queue.moveAsValue();
      }
    };

    return Base::customVisit(Visitor{val});
  }

 private:
  using Base::assignOr;
  using Base::data_;
};

} // namespace detail
} // namespace op
} // namespace thrift
} // namespace apache
