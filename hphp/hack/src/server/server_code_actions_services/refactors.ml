(*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)
let find ~entry ~(range : Lsp.range) ctx =
  let variable_actions =
    match Inline_variable.find ~entry ~range ctx with
    | [] -> Extract_variable.find ~entry ~range ctx
    | actions -> actions
  in
  Override_method.find ~entry ~range ctx
  @ variable_actions
  @ Extract_method.find ~entry ~range ctx
  @ Flip_around_comma.find ~entry ~range ctx
