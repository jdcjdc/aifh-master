// Artificial Intelligence for Humans
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
namespace AIFH_Vol3.Core
{
    /// <summary>
    /// Global constants for AIFH.
    /// </summary>
    public class AIFH
    {
        /// <summary>
        /// The default precision.
        /// </summary>
        public const double DefaultPrecision = 0.0000001;

        /// <summary>
        /// Private constructor.
        /// </summary>
        private AIFH()
        {

        }

        /// <summary>
        /// Allocate a 2D array.
        /// </summary>
        /// <typeparam name="T">The type to allocate.</typeparam>
        /// <param name="rows">Rows</param>
        /// <param name="cols">Columns</param>
        /// <returns>The array.</returns>
        public static T[][] Alloc2D<T>(int rows, int cols)
        {
            T[][] result = new T[rows][];
            for (int i = 0; i < rows; i++)
            {
                result[i] = new T[cols];
            }
            return result;
        }
    }
}
