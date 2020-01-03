﻿// Artificial Intelligence for Humans
// Volume 3: Deep Learning and Neural Networks
// C# Version
// http://www.aifh.org
// http://www.jeffheaton.com
//
// Code repository:
// https://github.com/jeffheaton/aifh
//
// Copyright 2015 by Jeff Heaton
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// For more information on Heaton Research copyrights, licenses
// and trademarks visit:
// http://www.heatonresearch.com/copyright
//

using AIFH_Vol3_Core.Core.ANN.Activation;

namespace AIFH_Vol3_Core.Core.ANN.Train.Error
{
    /// <summary>
    ///     Implements a cross entropy error function.  This can be used with backpropagation to
    ///     sometimes provide better performance than the standard linear error function.
    ///     De Boer, Pieter-Tjerk, et al. "A tutorial on the cross-entropy method." Annals of operations
    ///     research 134.1 (2005): 19-67.
    /// </summary>
    public class CrossEntropyErrorFunction : IErrorFunction
    {
        /// <inheritdoc />
        public void CalculateError(IActivationFunction af, double[] b, double[] a,
            double[] ideal, double[] actual, double[] error, double derivShift,
            double significance)
        {
            for (var i = 0; i < actual.Length; i++)
            {
                error[i] = (ideal[i] - actual[i])*significance;
            }
        }
    }
}